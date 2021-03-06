package gc_test

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
	"time"

	"code.cloudfoundry.org/lager/lagertest"

	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db/lock"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/postgresrunner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/ifrit"

	"testing"
)

func TestGc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gc Suite")
}

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

var (
	postgresRunner postgresrunner.Runner
	dbProcess      ifrit.Process

	dbConn                dbng.Conn
	err                   error
	resourceCacheFactory  dbng.ResourceCacheFactory
	resourceConfigFactory dbng.ResourceConfigFactory
	buildFactory          dbng.BuildFactory

	teamFactory dbng.TeamFactory

	defaultTeam     dbng.Team
	defaultPipeline dbng.Pipeline
	defaultBuild    dbng.Build

	usedResource dbng.Resource
	logger       *lagertest.TestLogger
)

var _ = BeforeSuite(func() {
	postgresRunner = postgresrunner.Runner{
		Port: 5433 + GinkgoParallelNode(),
	}

	dbProcess = ifrit.Invoke(postgresRunner)

	postgresRunner.CreateTestDB()
})

var _ = BeforeEach(func() {
	postgresRunner.Truncate()

	dbConn = postgresRunner.OpenConn()

	lockFactory := lock.NewLockFactory(postgresRunner.OpenSingleton())

	eBlock, err := aes.NewCipher([]byte("AES256Key-32Characters1234567890"))
	Expect(err).ToNot(HaveOccurred())
	aesgcm, err := cipher.NewGCM(eBlock)
	Expect(err).ToNot(HaveOccurred())
	key := dbng.NewEncryptionKey(aesgcm)

	teamFactory = dbng.NewTeamFactory(dbConn, lockFactory, key)
	buildFactory = dbng.NewBuildFactory(dbConn, lockFactory, key)

	defaultTeam, err = teamFactory.CreateTeam(atc.Team{Name: "default-team"})
	Expect(err).NotTo(HaveOccurred())

	defaultBuild, err = defaultTeam.CreateOneOffBuild()
	Expect(err).NotTo(HaveOccurred())

	atcConfig := atc.Config{
		Resources: atc.ResourceConfigs{
			{
				Name:   "some-resource",
				Type:   "resource-type",
				Source: atc.Source{"some": "source"},
			},
		},
	}

	defaultPipeline, _, err = defaultTeam.SavePipeline("default-pipeline", atcConfig, dbng.ConfigVersion(0), dbng.PipelineUnpaused)
	Expect(err).NotTo(HaveOccurred())

	var found bool
	usedResource, found, err = defaultPipeline.Resource("some-resource")
	Expect(err).NotTo(HaveOccurred())
	Expect(found).To(BeTrue())

	setupTx, err := dbConn.Begin()
	Expect(err).ToNot(HaveOccurred())

	baseResourceType := dbng.BaseResourceType{
		Name: "some-base-type",
	}
	_, err = baseResourceType.FindOrCreate(setupTx)
	Expect(err).NotTo(HaveOccurred())

	Expect(setupTx.Commit()).To(Succeed())

	logger = lagertest.NewTestLogger("gc-test")

	resourceCacheFactory = dbng.NewResourceCacheFactory(dbConn, lockFactory)
	resourceConfigFactory = dbng.NewResourceConfigFactory(dbConn, lockFactory)
})

var _ = AfterEach(func() {
	Expect(dbConn.Close()).To(Succeed())
})

var _ = AfterSuite(func() {
	dbProcess.Signal(os.Interrupt)
	Eventually(dbProcess.Wait(), 10*time.Second).Should(Receive())
})

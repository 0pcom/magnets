/*sitemap.go*/
package sitemap

import (
	//"github.com/sirupsen/logrus"
	//"github.com/skycoin/skycoin/src/util/logging"
	"github.com/spf13/cobra"
	sitemap "github.com/0pcom/magnets/pkg/sitemap"
	"strconv"
	"log"
	"os"
	"path"
	"sort"
	"time"
)


var (
	webPort	int
	webPort1 int
	host      string
	urlLoc    string
	urlLoc1    string
	outputDir string
	outputDir1 string
)

func init() {

}

func init() {
	wordDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		wordDir = ""
	}
	wordDir1 := wordDir
	wordDir = wordDir + "/public"
	wordDir1 = wordDir1 + "/public1"
	RootCmd.Flags().IntVarP(&webPort, "port", "p", 8040, "the host port")
	RootCmd.Flags().IntVarP(&webPort1, "port1", "q", 8041, "secondary host port")
	RootCmd.Flags().StringVarP(&host, "host", "i", "http://localhost", "the host name")
	RootCmd.Flags().StringVarP(&urlLoc, "loc", "l", "https://magnetosphere.net", "the prefix of sitemap loc tags")
	RootCmd.Flags().StringVarP(&urlLoc, "loc1", "m", "https://magnetosphereelectronicsurplus.com", "secondary prefix of sitemap loc tags")
	RootCmd.Flags().StringVarP(&outputDir, "out", "n", wordDir, "the sitemap output dir")
	RootCmd.Flags().StringVarP(&outputDir1, "out1", "o", wordDir1, "secondary sitemap output dir")
}

var RootCmd = &cobra.Command{
	Use:   "sitemap",
	Short: "generate a sitemap for the running web application",
//	PreRun: func(_ *cobra.Command, _ []string) {
//		wordDir, err := os.Getwd()
//		if err != nil {
//			log.Println(err)
//			wordDir = ""
//		}	},
	Run: func(_ *cobra.Command, _ []string) {
		//mLog := logging.NewMasterLogger()
		//mLog.SetLevel(logrus.InfoLevel)
		start := time.Now()
		origin := host
		origin1 := host
		wp := strconv.Itoa(webPort)
		wp1 := strconv.Itoa(webPort1)
		if wp != "80" {	origin += ":" + wp }
		if wp1 != "80" {	origin1 += ":" + wp1 }
		if urlLoc == "" {	urlLoc = origin	}
		if urlLoc1 == "" {	urlLoc1 = origin1	}
		s := &sitemap.Sitemap{}
		s1 := &sitemap.Sitemap{}
		s.Filename = path.Join(outputDir, "sitemap.xml")
		s1.Filename = path.Join(outputDir1, "sitemap.xml")
		s.Path = urlLoc
		s1.Path = urlLoc1
		log.Println("Crawling Host:", origin)
		log.Println("Crawling Host:", origin1)
		log.Println("Urlset Loc:", s.Path)
		log.Println("Urlset Loc:", s1.Path)
		log.Println("Sitemap File:", s.Filename)
		log.Println("Sitemap File:", s1.Filename)
		// Start crawling on the home page
		c := sitemap.NewCrawler(origin)
		c1 := sitemap.NewCrawler(origin1)
		cs := sitemap.NewCrawlerSupervisor(c)
		cs1 := sitemap.NewCrawlerSupervisor(c1)
		cs.AddJobToBuffer("/")
		cs1.AddJobToBuffer("/")
		// Block main until the crawler is done
		done := make(chan bool, 1)
		cs.Start(done)
		<-done
		close(done)
		done1 := make(chan bool, 1)
		cs1.Start(done1)
		<-done1
		s.Links = cs.GetVisitedLinks()
		s1.Links = cs1.GetVisitedLinks()
		sort.Strings(s.Links)
		sort.Strings(s1.Links)
		if err := s.WriteToFile(); err != nil {	log.Println(err)	}
		if err := s1.WriteToFile(); err != nil {	log.Println(err)	}
		log.Println("Finished in", time.Since(start))

	},
}

package main

import(
  "github.com/gin-gonic/gin"
)

func GetDummyEndpoint(c *gin.Context) {
  resp := map[string]string{"hello":"world"}
  c.JSON(200, resp)
}

func main() {
  api := gin.Default()
  api.GET("/dummy", GetDummyEndpoint)
  api.Run(":5000")
}
func RevisionMiddleware() gin.HandlerFunc {
  
  data, err := ioutil.ReadFile("REVISION")

  
  if err != nil {
    
    log.Println("revision middleware error:", err)

    return func(c *gin.Context) {
      c.Next()
    }
  }

  
  revision := strings.TrimSpace(string(data))

  
  return func(c *gin.Context) {
    c.Writer.Header().Set("X-Revision", revision)
    c.Next()
  }
}
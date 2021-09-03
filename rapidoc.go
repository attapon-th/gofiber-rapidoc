package rapidoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"html/template"

	"github.com/gofiber/fiber/v2"
)

type Config = RapiDocConfig

func New(config ...Config) fiber.Handler {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	} else {
		cfg = GetDefaultRapiDocConfig()
	}

	if rc, err := json.Marshal(&cfg); err != nil {
		panic(fmt.Errorf("OpenAPIURL is not empty."))
	} else {
		cfg = GetDefaultRapiDocConfig()
		json.Unmarshal(rc, &cfg)
	}

	return func(c *fiber.Ctx) error {
		p := c.Params("*", "index.html")
		if p == "index.html" {
			tpl, err := template.New("rapidoc").Parse(HtmlTemplateRapiDoc())
			if err != nil {
				return err
			}

			buf := new(bytes.Buffer)
			// log.Println(cfg)
			if err := tpl.Execute(buf, cfg); err != nil {
				return err
			}
			c.Set("content-type", "text/html")
			return c.SendStream(buf)

		} else if strings.HasSuffix(p, ".json") || strings.HasSuffix(p, ".yaml") || strings.HasSuffix(p, ".yml") {
			if _, err := os.Stat(p); err == nil {
				if strings.HasSuffix(p, ".json") {
					c.Set("content-type", "application/json")
				} else {
					c.Set("content-type", "text/yaml")
				}
				return c.SendFile(p, true)
			}
		}
		return c.SendStatus(404)
	}

}

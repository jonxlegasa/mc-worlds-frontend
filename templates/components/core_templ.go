// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/templates/helpers"
	"strings"
)

func Metatags(page *controller.Page) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(page.AppName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 12, Col: 16}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(page.Title) > 0 {
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("| %s", page.Title))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 14, Col: 36}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><link rel=\"icon\" href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.File("favicon.png"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 17, Col: 52}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"viewport-fit=cover, width=device-width, initial-scale=1.0, shrink-to-fit=no\"><meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(page.Metatags.Description) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<meta name=\"description\" content=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(page.Metatags.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 22, Col: 62}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if len(page.Metatags.Keywords) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<meta name=\"keywords\" content=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(strings.Join(page.Metatags.Keywords, ", "))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 25, Col: 76}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<meta name=\"application-name\" content=\"Goship\"><meta name=\"msapplication-starturl\" content=\"https://cherie.chatbond.app\"><meta name=\"mobile-web-app-capable\" content=\"yes\"><meta name=\"apple-mobile-web-app-capable\" content=\"yes\"><meta name=\"apple-mobile-web-app-status-bar-style\" content=\"black-translucent\"><meta name=\"apple-mobile-web-app-title\" content=\"Goship\"><link rel=\"apple-touch-startup-image\" href=\"ios-startup.png\"><meta name=\"color-scheme\" content=\"light dark\"><meta name=\"description\" content=\"Goship is a Go/HTMX starter kit here to help you get your idea to market before you can sneeze once.\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func CSS() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<link rel=\"stylesheet\" href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.File("svelte_bundle.css"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 39, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><link rel=\"stylesheet\" href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.File("styles_bundle.css"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 40, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><link rel=\"manifest\" crossorigin=\"use-credentials\" href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.File("manifest.json"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 41, Col: 88}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><link rel=\"preload\" href=\"https://cdn.jsdelivr.net/npm/swiper@11/swiper-bundle.min.css\" as=\"style\" onload=\"this.onload=null;this.rel=&#39;stylesheet&#39;\"><noscript><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/swiper@11/swiper-bundle.min.css\"></noscript><link rel=\"preload\" href=\"https://unpkg.com/tippy.js@6/dist/tippy.css\" as=\"style\" onload=\"this.onload=null;this.rel=&#39;stylesheet&#39;\"><noscript><link rel=\"stylesheet\" href=\"https://unpkg.com/tippy.js@6/dist/tippy.css\"></noscript><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/driver.js@1.0.1/dist/driver.css\"><style>\n\tbody {\n\t\tbackground-color: transparent;\n\t}\n\n\t:root {\n\t\t--safe-area-top: env(safe-area-inset-top);\n\t\t--safe-area-left: env(safe-area-inset-left);\n\t}\n\n\t#navbar {\n\t\tpadding-top: env(safe-area-inset-top);\n\t\t// padding-left: env(safe-area-inset-left);\n\t\tpadding-right: 0; /* Adjust as needed */\n\t\tpadding-bottom: 0; /* Adjust as needed */\n\t}\n\t</style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func JS() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var11 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var11 == nil {
			templ_7745c5c3_Var11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"https://cdn.jsdelivr.net/npm/ios-pwa-splash@1.0.0/cdn.min.js\"></script><script id=\"dynamic-splash-screen\" icon-url-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(helpers.File("icon.png")))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 74, Col: 94}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">\n\t\t(function(){\n\t\t\tconst dynamicSplashScreen = document.getElementById(\"dynamic-splash-screen\");\n\t\t\tconst iconURL = JSON.parse(dynamicSplashScreen.getAttribute(\"icon-url-data\"));\n\t\t\tconsole.log(`setting up iosPWASpash with icon url ${iconURL}`)\n\t\t\tiosPWASplash(iconURL, \"#000000\");\n\t\t})()\n\t</script><script src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.File("main.js"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 83, Col: 38}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></script><script src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.File("vanilla_bundle.js"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 84, Col: 48}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></script><script src=\"https://cdn.jsdelivr.net/npm/sweetalert2@10\"></script><script src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(helpers.ServiceWorkerFile("service-worker.js"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 87, Col: 61}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></script><script src=\"https://unpkg.com/htmx.org@1.9.10\" integrity=\"sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC\" crossorigin=\"anonymous\"></script><script src=\"https://unpkg.com/htmx.org@1.9.12/dist/ext/sse.js\"></script><script defer async src=\"https://cdn.jsdelivr.net/npm/@alpinejs/mask@3.x.x/dist/cdn.min.js\"></script><script defer async src=\"https://cdn.jsdelivr.net/npm/@ryangjchandler/alpine-clipboard@2.x.x/dist/alpine-clipboard.js\" defer></script><script defer src=\"https://cdn.jsdelivr.net/npm/@alpinejs/collapse@3.x.x/dist/cdn.min.js\"></script><script defer async src=\"https://cdn.jsdelivr.net/npm/@marcreichel/alpine-timeago@latest/dist/alpine-timeago.min.js\" defer></script><script defer async src=\"https://cdn.jsdelivr.net/npm/@ryangjchandler/alpine-tooltip@1.x.x/dist/cdn.min.js\" defer></script><script defer async src=\"https://cdn.jsdelivr.net/npm/@alpinejs/morph@3.x.x/dist/cdn.min.js\"></script><script defer async src=\"https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js\"></script><script defer async src=\"https://unpkg.com/hyperscript.org@0.9.12\"></script><script defer async src=\"https://cdn.jsdelivr.net/npm/swiper@11/swiper-bundle.min.js\"></script><script defer async src=\"https://js.stripe.com/v3/\"></script><script defer async src=\"https://cdnjs.cloudflare.com/ajax/libs/flowbite/2.3.0/datepicker.min.js\"></script><script defer async src=\"https://cdn.jsdelivr.net/npm/dayjs@1/dayjs.min.js\"></script><script defer async src=\"https://d3js.org/d3.v7.min.js\"></script><script defer async src=\"https://unpkg.com/cal-heatmap/dist/cal-heatmap.min.js\"></script><script defer async src=\"https://unpkg.com/cal-heatmap/dist/plugins/Legend.min.js\"></script><script defer async src=\"https://unpkg.com/cal-heatmap/dist/plugins/Tooltip.min.js\"></script><script defer async src=\"https://unpkg.com/cal-heatmap/dist/plugins/CalendarLabel.min.js\"></script><script defer async src=\"https://unpkg.com/@popperjs/core@2\"></script><script src=\"https://cdn.jsdelivr.net/npm/driver.js@1.0.1/dist/driver.js.iife.js\"></script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = darkModeSwitcher().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func sentrySessionReplay() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_sentrySessionReplay_e338`,
		Function: `function __templ_sentrySessionReplay_e338(){<script defer async 
	src="https://js.sentry-cdn.com/1a4df3451703a57a67ad24c6c517a50c.min.js"
	crossorigin="anonymous"
	></script>
}`,
		Call:       templ.SafeScript(`__templ_sentrySessionReplay_e338`),
		CallInline: templ.SafeScriptInline(`__templ_sentrySessionReplay_e338`),
	}
}

func darkModeSwitcher() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_darkModeSwitcher_5a44`,
		Function: `function __templ_darkModeSwitcher_5a44(){// On page load or when changing themes, best to add inline in ` + "`" + `head` + "`" + ` to avoid FOUC
    if (localStorage.getItem('color-theme') === 'darkmode' || (!('color-theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        document.documentElement.classList.add('dark'); // For default tailwind dark: utility
		document.documentElement.setAttribute('data-theme', 'darkmode'); // For daisyui theme
		// Set the hover brightness dynamically
		document.documentElement.style.setProperty('--brightness-hover', 'var(--brightness-hover-dark)');
    } else {
        document.documentElement.classList.remove('dark') // For default tailwind dark: utility
		document.documentElement.setAttribute('data-theme', 'lightmode'); // For daisyui theme
		// Set the hover brightness dynamically
		document.documentElement.style.setProperty('--brightness-hover', 'var(--brightness-hover-light)');
    }
}`,
		Call:       templ.SafeScript(`__templ_darkModeSwitcher_5a44`),
		CallInline: templ.SafeScriptInline(`__templ_darkModeSwitcher_5a44`),
	}
}

func csrfJS(token string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_csrfJS_e551`,
		Function: `function __templ_csrfJS_e551(token){document.body.addEventListener('htmx:configRequest', function(evt)  {
        if (evt.detail.verb !== "get") {
            evt.detail.parameters['csrf'] = token;
        }
    })
}`,
		Call:       templ.SafeScript(`__templ_csrfJS_e551`, token),
		CallInline: templ.SafeScriptInline(`__templ_csrfJS_e551`, token),
	}
}

func htmxBeforeSwap() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_htmxBeforeSwap_7d49`,
		Function: `function __templ_htmxBeforeSwap_7d49(){document.body.addEventListener('htmx:beforeSwap', function(evt) {
        if (evt.detail.xhr.status >= 400){
            evt.detail.shouldSwap = true;
            evt.detail.target = htmx.find("body");
        }
    });
}`,
		Call:       templ.SafeScript(`__templ_htmxBeforeSwap_7d49`),
		CallInline: templ.SafeScriptInline(`__templ_htmxBeforeSwap_7d49`),
	}
}

func JSFooter(page *controller.Page) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var16 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var16 == nil {
			templ_7745c5c3_Var16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if len(page.CSRF) > 0 {
			templ_7745c5c3_Err = csrfJS(page.CSRF).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = htmxBeforeSwap().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = htmxOnLoad().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = beforeBodyEnd().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n\t\tfunction handleCacheUpdate(event) {\n\t\t\tconst url = new URL(event.detail.requestConfig.url, location.origin);\n\t\t\tif (navigator.serviceWorker && navigator.serviceWorker.controller) {\n\t\t\t\tnavigator.serviceWorker.controller.postMessage({\n\t\t\t\ttype: \"CACHE_UPDATED\",\n\t\t\t\turl: url.href,\n\t\t\t\t});\n\t\t\t}\n\t\t}\n\t\tfunction handleCacheUpdateWithURL(url) {\n\t\t\tif (navigator.serviceWorker && navigator.serviceWorker.controller) {\n\t\t\tnavigator.serviceWorker.controller.postMessage({\n\t\t\t\ttype: 'CACHE_UPDATE',\n\t\t\t\turl: url\n\t\t\t});\n\t\t\t}\n\t\t}\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func htmxOnLoad() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_htmxOnLoad_aeb5`,
		Function: `function __templ_htmxOnLoad_aeb5(){htmx.onLoad(function(content) {
		initializeJS();
	});	
	
}`,
		Call:       templ.SafeScript(`__templ_htmxOnLoad_aeb5`),
		CallInline: templ.SafeScriptInline(`__templ_htmxOnLoad_aeb5`),
	}
}

func htmxWindowOnLoad() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_htmxWindowOnLoad_6a72`,
		Function: `function __templ_htmxWindowOnLoad_6a72(){// Was trying to get CTRL/CMD + Click to work, but still not working. HTMX uses hx-get instead of href,
// which confuses browsers. https://www.joncom.be/notes/htmx-ctrl-click/

	window.onload = function () {
		htmx.on{'htmx:configRequest', (evt) => {
			const requestType = evt.detail.verb;
			const ctrlKey = evt.detail.triggeringEvent?.ctrlKey;
			const pushUrl = evt.target.attributes['hx-push-url']?.value;

			if (requestType === 'get' && ctrlKey && pushUrl) {
				evt.preventDefault();
				window.open(pushUrl, '_blank');
			}
		}
		}
	}
}`,
		Call:       templ.SafeScript(`__templ_htmxWindowOnLoad_6a72`),
		CallInline: templ.SafeScriptInline(`__templ_htmxWindowOnLoad_6a72`),
	}
}

func beforeBodyEnd() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var17 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var17 == nil {
			templ_7745c5c3_Var17 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/flowbite/2.2.1/flowbite.min.js\"></script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func TextFooter(page *controller.Page) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var18 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var18 == nil {
			templ_7745c5c3_Var18 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mx-1 sm:mx-2 py-3\"><footer class=\"bg-slate-100 dark:bg-gray-800 rounded-lg shadow m-4\"><div class=\"w-full mx-auto max-w-screen-xl p-4 md:flex md:items-center md:justify-between\"><span class=\"text-sm text-gray-500 sm:text-center dark:text-gray-400\">© <span id=\"currentYear\"></span> <a href=\"https://cherie.chatbond.app\" class=\"hover:underline\">Chatbond Inc</a>. All Rights Reserved.</span><ul hx-target=\"#main-content\" hx-select=\"#main-content\" hx-indicator=\"next #page-loading\" hx-swap=\"outerHTML show:window:top\" hx-push-url=\"true\" class=\"flex flex-wrap items-center mt-3 text-sm font-medium text-gray-500 dark:text-gray-400 sm:mt-0\"><li><a hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var19 string
		templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(page.ToURL("about"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 235, Col: 35}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"hover:underline me-4 md:me-6 cursor-pointer\">Contact Us</a></li><li><a hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var20 string
		templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(page.ToURL("privacy_policy"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/core.templ`, Line: 243, Col: 44}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"hover:underline me-4 md:me-6 cursor-pointer\">Privacy Policy</a></li></ul></div></footer></div><script>\n\t\tdocument.getElementById('currentYear').textContent = new Date().getFullYear();\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate

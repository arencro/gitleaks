package rules

// TODO introduce skiplists:
// https://github.com/danielmiessler/SecLists/blob/master/Miscellaneous/wordlist-skipfish.fuzz.txt
// https://github.com/e3b0c442/keywords
// https://gist.github.com/maxtruxa/b2ca551e42d3aead2b3d
// https://github.com/HChakraborty/projects/commit/e860cb863ee9585c38db8360814b04ef9fa1bdce
// https://github.com/UraniumX92/Discord-Bot-using-py/tree/224b2b71a58c25f420ce980f2ea49627b4b646f1/Data%20Files
// https://github.com/Meen11/BSBI-Indexing/blob/63032017aa24f3111f18468607cd0db5997bb891/datasets/citeseer/11/10.1.1.27.6385.txt

var DefaultStopWords = []string{
	"client",
	"endpoint",
	"vpn",
	"_ec2_",
	"aws_",
	"authorize",
	"author",
	"define",
	"config",
	"credential",
	"setting",
	"sample",
	"xxxxxx",
	"000000",
	"buffer",
	"delete",
	"aaaaaa",
	"fewfwef",
	"getenv",
	"env_",
	"system",
	"example",
	"ecdsa",
	"sha256",
	"sha1",
	"sha2",
	"md5",
	"alert",
	"wizard",
	"target",
	"onboard",
	"welcome",
	"page",
	"exploit",
	"experiment",
	"expire",
	"rabbitmq",
	"scraper",
	"widget",
	"music",
	"dns_",
	"dns-",
	"yahoo",
	"want",
	"json",
	"action",
	"script",
	"fix_",
	"fix-",
	"develop",
	"compas",
	"stripe",
	"service",
	"master",
	"metric",
	"tech",
	"gitignore",
	"rich",
	"open",
	"stack",
	"irc_",
	"irc-",
	"sublime",
	"kohana",
	"has_",
	"has-",
	"fabric",
	"wordpres",
	"role",
	"osx_",
	"osx-",
	"boost",
	"addres",
	"queue",
	"working",
	"sandbox",
	"internet",
	"print",
	"vision",
	"tracking",
	"being",
	"generator",
	"traffic",
	"world",
	"pull",
	"rust",
	"watcher",
	"small",
	"auth",
	"full",
	"hash",
	"more",
	"install",
	"auto",
	"complete",
	"learn",
	"paper",
	"installer",
	"research",
	"acces",
	"last",
	"binding",
	"spine",
	"into",
	"chat",
	"algorithm",
	"resource",
	"uploader",
	"video",
	"maker",
	"next",
	"proc",
	"lock",
	"robot",
	"snake",
	"patch",
	"matrix",
	"drill",
	"terminal",
	"term",
	"stuff",
	"genetic",
	"generic",
	"identity",
	"audit",
	"pattern",
	"audio",
	"web_",
	"web-",
	"crud",
	"problem",
	"statu",
	"cms-",
	"cms_",
	"arch",
	"coffee",
	"workflow",
	"changelog",
	"another",
	"uiview",
	"content",
	"kitchen",
	"gnu_",
	"gnu-",
	"gnu.",
	"conf",
	"couchdb",
	"client",
	"opencv",
	"rendering",
	"update",
	"concept",
	"varnish",
	"gui_",
	"gui-",
	"gui.",
	"version",
	"shared",
	"extra",
	"product",
	"still",
	"not_",
	"not-",
	"not.",
	"drop",
	"ring",
	"png_",
	"png-",
	"png.",
	"actively",
	"import",
	"output",
	"backup",
	"start",
	"embedded",
	"registry",
	"pool",
	"semantic",
	"instagram",
	"bash",
	"system",
	"ninja",
	"drupal",
	"jquery",
	"polyfill",
	"physic",
	"league",
	"guide",
	"pack",
	"synopsi",
	"sketch",
	"injection",
	"svg_",
	"svg-",
	"svg.",
	"friendly",
	"wave",
	"convert",
	"manage",
	"camera",
	"link",
	"slide",
	"timer",
	"wrapper",
	"gallery",
	"url_",
	"url-",
	"url.",
	"todomvc",
	"requirej",
	"party",
	"http",
	"payment",
	"async",
	"library",
	"home",
	"coco",
	"gaia",
	"display",
	"universal",
	"function",
	"metadata",
	"hipchat",
	"under",
	"room",
	"config",
	"personal",
	"realtime",
	"resume",
	"database",
	"testing",
	"tiny",
	"basic",
	"forum",
	"meetup",
	"yet_",
	"yet-",
	"yet.",
	"cento",
	"dead",
	"fluentd",
	"editor",
	"utilitie",
	"run_",
	"run-",
	"run.",
	"box_",
	"box-",
	"box.",
	"bot_",
	"bot-",
	"bot.",
	"making",
	"sample",
	"group",
	"monitor",
	"ajax",
	"parallel",
	"cassandra",
	"ultimate",
	"site",
	"get_",
	"get-",
	"get.",
	"gen_",
	"gen-",
	"gen.",
	"gem_",
	"gem-",
	"gem.",
	"extended",
	"image",
	"knife",
	"asset",
	"nested",
	"zero",
	"plugin",
	"bracket",
	"mule",
	"mozilla",
	"number",
	"act_",
	"act-",
	"act.",
	"map_",
	"map-",
	"map.",
	"micro",
	"debug",
	"openshift",
	"chart",
	"expres",
	"backend",
	"task",
	"source",
	"translate",
	"jbos",
	"composer",
	"sqlite",
	"profile",
	"mustache",
	"mqtt",
	"yeoman",
	"have",
	"builder",
	"smart",
	"like",
	"oauth",
	"school",
	"guideline",
	"captcha",
	"filter",
	"bitcoin",
	"bridge",
	"color",
	"toolbox",
	"discovery",
	"new_",
	"new-",
	"new.",
	"dashboard",
	"when",
	"setting",
	"level",
	"post",
	"standard",
	"port",
	"platform",
	"yui_",
	"yui-",
	"yui.",
	"grunt",
	"animation",
	"haskell",
	"icon",
	"latex",
	"cheat",
	"lua_",
	"lua-",
	"lua.",
	"gulp",
	"case",
	"author",
	"without",
	"simulator",
	"wifi",
	"directory",
	"lisp",
	"list",
	"flat",
	"adventure",
	"story",
	"storm",
	"gpu_",
	"gpu-",
	"gpu.",
	"store",
	"caching",
	"attention",
	"solr",
	"logger",
	"demo",
	"shortener",
	"hadoop",
	"finder",
	"phone",
	"pipeline",
	"range",
	"textmate",
	"showcase",
	"app_",
	"app-",
	"app.",
	"idiomatic",
	"edit",
	"our_",
	"our-",
	"our.",
	"out_",
	"out-",
	"out.",
	"sentiment",
	"linked",
	"why_",
	"why-",
	"why.",
	"local",
	"cube",
	"gmail",
	"job_",
	"job-",
	"job.",
	"rpc_",
	"rpc-",
	"rpc.",
	"contest",
	"tcp_",
	"tcp-",
	"tcp.",
	"usage",
	"buildout",
	"weather",
	"transfer",
	"automated",
	"sphinx",
	"issue",
	"sas_",
	"sas-",
	"sas.",
	"parallax",
	"jasmine",
	"addon",
	"machine",
	"solution",
	"dsl_",
	"dsl-",
	"dsl.",
	"episode",
	"menu",
	"theme",
	"best",
	"adapter",
	"debugger",
	"chrome",
	"tutorial",
	"life",
	"step",
	"people",
	"joomla",
	"paypal",
	"developer",
	"solver",
	"team",
	"current",
	"love",
	"visual",
	"date",
	"data",
	"canva",
	"container",
	"future",
	"xml_",
	"xml-",
	"xml.",
	"twig",
	"nagio",
	"spatial",
	"original",
	"sync",
	"archived",
	"refinery",
	"science",
	"mapping",
	"gitlab",
	"play",
	"ext_",
	"ext-",
	"ext.",
	"session",
	"impact",
	"set_",
	"set-",
	"set.",
	"see_",
	"see-",
	"see.",
	"migration",
	"commit",
	"community",
	"shopify",
	"what'",
	"cucumber",
	"statamic",
	"mysql",
	"location",
	"tower",
	"line",
	"code",
	"amqp",
	"hello",
	"send",
	"index",
	"high",
	"notebook",
	"alloy",
	"python",
	"field",
	"document",
	"soap",
	"edition",
	"email",
	"php_",
	"php-",
	"php.",
	"command",
	"transport",
	"official",
	"upload",
	"study",
	"secure",
	"angularj",
	"akka",
	"scalable",
	"package",
	"request",
	"con_",
	"con-",
	"con.",
	"flexible",
	"security",
	"comment",
	"module",
	"flask",
	"graph",
	"flash",
	"apache",
	"change",
	"window",
	"space",
	"lambda",
	"sheet",
	"bookmark",
	"carousel",
	"friend",
	"objective",
	"jekyll",
	"bootstrap",
	"first",
	"article",
	"gwt_",
	"gwt-",
	"gwt.",
	"classic",
	"media",
	"websocket",
	"touch",
	"desktop",
	"real",
	"read",
	"recorder",
	"moved",
	"storage",
	"validator",
	"add-on",
	"pusher",
	"scs_",
	"scs-",
	"scs.",
	"inline",
	"asp_",
	"asp-",
	"asp.",
	"timeline",
	"base",
	"encoding",
	"ffmpeg",
	"kindle",
	"tinymce",
	"pretty",
	"jpa_",
	"jpa-",
	"jpa.",
	"used",
	"user",
	"required",
	"webhook",
	"download",
	"resque",
	"espresso",
	"cloud",
	"mongo",
	"benchmark",
	"pure",
	"cakephp",
	"modx",
	"mode",
	"reactive",
	"fuel",
	"written",
	"flickr",
	"mail",
	"brunch",
	"meteor",
	"dynamic",
	"neo_",
	"neo-",
	"neo.",
	"new_",
	"new-",
	"new.",
	"net_",
	"net-",
	"net.",
	"typo",
	"type",
	"keyboard",
	"erlang",
	"adobe",
	"logging",
	"ckeditor",
	"message",
	"iso_",
	"iso-",
	"iso.",
	"hook",
	"ldap",
	"folder",
	"reference",
	"railscast",
	"www_",
	"www-",
	"www.",
	"tracker",
	"azure",
	"fork",
	"form",
	"digital",
	"exporter",
	"skin",
	"string",
	"template",
	"designer",
	"gollum",
	"fluent",
	"entity",
	"language",
	"alfred",
	"summary",
	"wiki",
	"kernel",
	"calendar",
	"plupload",
	"symfony",
	"foundry",
	"remote",
	"talk",
	"search",
	"dev_",
	"dev-",
	"dev.",
	"del_",
	"del-",
	"del.",
	"token",
	"idea",
	"sencha",
	"selector",
	"interface",
	"create",
	"fun_",
	"fun-",
	"fun.",
	"groovy",
	"query",
	"grail",
	"red_",
	"red-",
	"red.",
	"laravel",
	"monkey",
	"slack",
	"supported",
	"instant",
	"value",
	"center",
	"latest",
	"work",
	"but_",
	"but-",
	"but.",
	"bug_",
	"bug-",
	"bug.",
	"virtual",
	"tweet",
	"statsd",
	"studio",
	"path",
	"real-time",
	"frontend",
	"notifier",
	"coding",
	"tool",
	"firmware",
	"flow",
	"random",
	"mediawiki",
	"bosh",
	"been",
	"beer",
	"lightbox",
	"theory",
	"origin",
	"redmine",
	"hub_",
	"hub-",
	"hub.",
	"require",
	"pro_",
	"pro-",
	"pro.",
	"ant_",
	"ant-",
	"ant.",
	"any_",
	"any-",
	"any.",
	"recipe",
	"closure",
	"mapper",
	"event",
	"todo",
	"model",
	"redi",
	"provider",
	"rvm_",
	"rvm-",
	"rvm.",
	"program",
	"memcached",
	"rail",
	"silex",
	"foreman",
	"activity",
	"license",
	"strategy",
	"batch",
	"streaming",
	"fast",
	"use_",
	"use-",
	"use.",
	"usb_",
	"usb-",
	"usb.",
	"impres",
	"academy",
	"slider",
	"please",
	"layer",
	"cros",
	"now_",
	"now-",
	"now.",
	"miner",
	"extension",
	"own_",
	"own-",
	"own.",
	"app_",
	"app-",
	"app.",
	"debian",
	"symphony",
	"example",
	"feature",
	"serie",
	"tree",
	"project",
	"runner",
	"entry",
	"leetcode",
	"layout",
	"webrtc",
	"logic",
	"login",
	"worker",
	"toolkit",
	"mocha",
	"support",
	"back",
	"inside",
	"device",
	"jenkin",
	"contact",
	"fake",
	"awesome",
	"ocaml",
	"bit_",
	"bit-",
	"bit.",
	"drive",
	"screen",
	"prototype",
	"gist",
	"binary",
	"nosql",
	"rest",
	"overview",
	"dart",
	"dark",
	"emac",
	"mongoid",
	"solarized",
	"homepage",
	"emulator",
	"commander",
	"django",
	"yandex",
	"gradle",
	"xcode",
	"writer",
	"crm_",
	"crm-",
	"crm.",
	"jade",
	"startup",
	"error",
	"using",
	"format",
	"name",
	"spring",
	"parser",
	"scratch",
	"magic",
	"try_",
	"try-",
	"try.",
	"rack",
	"directive",
	"challenge",
	"slim",
	"counter",
	"element",
	"chosen",
	"doc_",
	"doc-",
	"doc.",
	"meta",
	"should",
	"button",
	"packet",
	"stream",
	"hardware",
	"android",
	"infinite",
	"password",
	"software",
	"ghost",
	"xamarin",
	"spec",
	"chef",
	"interview",
	"hubot",
	"mvc_",
	"mvc-",
	"mvc.",
	"exercise",
	"leaflet",
	"launcher",
	"air_",
	"air-",
	"air.",
	"photo",
	"board",
	"boxen",
	"way_",
	"way-",
	"way.",
	"computing",
	"welcome",
	"notepad",
	"portfolio",
	"cat_",
	"cat-",
	"cat.",
	"can_",
	"can-",
	"can.",
	"magento",
	"yaml",
	"domain",
	"card",
	"yii_",
	"yii-",
	"yii.",
	"checker",
	"browser",
	"upgrade",
	"only",
	"progres",
	"aura",
	"ruby_",
	"ruby-",
	"ruby.",
	"polymer",
	"util",
	"lite",
	"hackathon",
	"rule",
	"log_",
	"log-",
	"log.",
	"opengl",
	"stanford",
	"skeleton",
	"history",
	"inspector",
	"help",
	"soon",
	"selenium",
	"lab_",
	"lab-",
	"lab.",
	"scheme",
	"schema",
	"look",
	"ready",
	"leveldb",
	"docker",
	"game",
	"minimal",
	"logstash",
	"messaging",
	"within",
	"heroku",
	"mongodb",
	"kata",
	"suite",
	"picker",
	"win_",
	"win-",
	"win.",
	"wip_",
	"wip-",
	"wip.",
	"panel",
	"started",
	"starter",
	"front-end",
	"detector",
	"deploy",
	"editing",
	"based",
	"admin",
	"capture",
	"spree",
	"page",
	"bundle",
	"goal",
	"rpg_",
	"rpg-",
	"rpg.",
	"setup",
	"side",
	"mean",
	"reader",
	"cookbook",
	"mini",
	"modern",
	"seed",
	"dom_",
	"dom-",
	"dom.",
	"doc_",
	"doc-",
	"doc.",
	"dot_",
	"dot-",
	"dot.",
	"syntax",
	"sugar",
	"loader",
	"website",
	"make",
	"kit_",
	"kit-",
	"kit.",
	"protocol",
	"human",
	"daemon",
	"golang",
	"manager",
	"countdown",
	"connector",
	"swagger",
	"map_",
	"map-",
	"map.",
	"mac_",
	"mac-",
	"mac.",
	"man_",
	"man-",
	"man.",
	"orm_",
	"orm-",
	"orm.",
	"org_",
	"org-",
	"org.",
	"little",
	"zsh_",
	"zsh-",
	"zsh.",
	"shop",
	"show",
	"workshop",
	"money",
	"grid",
	"server",
	"octopres",
	"svn_",
	"svn-",
	"svn.",
	"ember",
	"embed",
	"general",
	"file",
	"important",
	"dropbox",
	"portable",
	"public",
	"docpad",
	"fish",
	"sbt_",
	"sbt-",
	"sbt.",
	"done",
	"para",
	"network",
	"common",
	"readme",
	"popup",
	"simple",
	"purpose",
	"mirror",
	"single",
	"cordova",
	"exchange",
	"object",
	"design",
	"gateway",
	"account",
	"lamp",
	"intellij",
	"math",
	"mit_",
	"mit-",
	"mit.",
	"control",
	"enhanced",
	"emitter",
	"multi",
	"add_",
	"add-",
	"add.",
	"about",
	"socket",
	"preview",
	"vagrant",
	"cli_",
	"cli-",
	"cli.",
	"powerful",
	"top_",
	"top-",
	"top.",
	"radio",
	"watch",
	"fluid",
	"amazon",
	"report",
	"couchbase",
	"automatic",
	"detection",
	"sprite",
	"pyramid",
	"portal",
	"advanced",
	"plu_",
	"plu-",
	"plu.",
	"runtime",
	"git_",
	"git-",
	"git.",
	"uri_",
	"uri-",
	"uri.",
	"haml",
	"node",
	"sql_",
	"sql-",
	"sql.",
	"cool",
	"core",
	"obsolete",
	"handler",
	"iphone",
	"extractor",
	"array",
	"copy",
	"nlp_",
	"nlp-",
	"nlp.",
	"reveal",
	"pop_",
	"pop-",
	"pop.",
	"engine",
	"parse",
	"check",
	"html",
	"nest",
	"all_",
	"all-",
	"all.",
	"chinese",
	"buildpack",
	"what",
	"tag_",
	"tag-",
	"tag.",
	"proxy",
	"style",
	"cookie",
	"feed",
	"restful",
	"compiler",
	"creating",
	"prelude",
	"context",
	"java",
	"rspec",
	"mock",
	"backbone",
	"light",
	"spotify",
	"flex",
	"related",
	"shell",
	"which",
	"clas",
	"webapp",
	"swift",
	"ansible",
	"unity",
	"console",
	"tumblr",
	"export",
	"campfire",
	"conway'",
	"made",
	"riak",
	"hero",
	"here",
	"unix",
	"unit",
	"glas",
	"smtp",
	"how_",
	"how-",
	"how.",
	"hot_",
	"hot-",
	"hot.",
	"debug",
	"release",
	"diff",
	"player",
	"easy",
	"right",
	"old_",
	"old-",
	"old.",
	"animate",
	"time",
	"push",
	"explorer",
	"course",
	"training",
	"nette",
	"router",
	"draft",
	"structure",
	"note",
	"salt",
	"where",
	"spark",
	"trello",
	"power",
	"method",
	"social",
	"via_",
	"via-",
	"via.",
	"vim_",
	"vim-",
	"vim.",
	"select",
	"webkit",
	"github",
	"ftp_",
	"ftp-",
	"ftp.",
	"creator",
	"mongoose",
	"led_",
	"led-",
	"led.",
	"movie",
	"currently",
	"pdf_",
	"pdf-",
	"pdf.",
	"load",
	"markdown",
	"phalcon",
	"input",
	"custom",
	"atom",
	"oracle",
	"phonegap",
	"ubuntu",
	"great",
	"rdf_",
	"rdf-",
	"rdf.",
	"popcorn",
	"firefox",
	"zip_",
	"zip-",
	"zip.",
	"cuda",
	"dotfile",
	"static",
	"openwrt",
	"viewer",
	"powered",
	"graphic",
	"les_",
	"les-",
	"les.",
	"doe_",
	"doe-",
	"doe.",
	"maven",
	"word",
	"eclipse",
	"lab_",
	"lab-",
	"lab.",
	"hacking",
	"steam",
	"analytic",
	"option",
	"abstract",
	"archive",
	"reality",
	"switcher",
	"club",
	"write",
	"kafka",
	"arduino",
	"angular",
	"online",
	"title",
	"don't",
	"contao",
	"notice",
	"analyzer",
	"learning",
	"zend",
	"external",
	"staging",
	"busines",
	"tdd_",
	"tdd-",
	"tdd.",
	"scanner",
	"building",
	"snippet",
	"modular",
	"bower",
	"stm_",
	"stm-",
	"stm.",
	"lib_",
	"lib-",
	"lib.",
	"alpha",
	"mobile",
	"clean",
	"linux",
	"nginx",
	"manifest",
	"some",
	"raspberry",
	"gnome",
	"ide_",
	"ide-",
	"ide.",
	"block",
	"statistic",
	"info",
	"drag",
	"youtube",
	"koan",
	"facebook",
	"paperclip",
	"art_",
	"art-",
	"art.",
	"quality",
	"tab_",
	"tab-",
	"tab.",
	"need",
	"dojo",
	"shield",
	"computer",
	"stat",
	"state",
	"twitter",
	"utility",
	"converter",
	"hosting",
	"devise",
	"liferay",
	"updated",
	"force",
	"tip_",
	"tip-",
	"tip.",
	"behavior",
	"active",
	"call",
	"answer",
	"deck",
	"better",
	"principle",
	"ches",
	"bar_",
	"bar-",
	"bar.",
	"reddit",
	"three",
	"haxe",
	"just",
	"plug-in",
	"agile",
	"manual",
	"tetri",
	"super",
	"beta",
	"parsing",
	"doctrine",
	"minecraft",
	"useful",
	"perl",
	"sharing",
	"agent",
	"switch",
	"view",
	"dash",
	"channel",
	"repo",
	"pebble",
	"profiler",
	"warning",
	"cluster",
	"running",
	"markup",
	"evented",
	"mod_",
	"mod-",
	"mod.",
	// "api_", lin_api_ is used for linear
	// "api-",
	// "api.",
	"share",
	"csv_",
	"csv-",
	"csv.",
	"response",
	"good",
	"house",
	"connect",
	"built",
	"build",
	"find",
	"ipython",
	"webgl",
	"big_",
	"big-",
	"big.",
	"google",
	"scala",
	"sdl_",
	"sdl-",
	"sdl.",
	"sdk_",
	"sdk-",
	"sdk.",
	"native",
	"day_",
	"day-",
	"day.",
	"puppet",
	"text",
	"routing",
	"helper",
	"linkedin",
	"crawler",
	"host",
	"guard",
	"merchant",
	"poker",
	"over",
	"writing",
	"free",
	"classe",
	"component",
	"craft",
	"nodej",
	"phoenix",
	"longer",
	"quick",
	"lazy",
	"memory",
	"clone",
	"hacker",
	"middleman",
	"factory",
	"motion",
	"multiple",
	"tornado",
	"hack",
	"ssh_",
	"ssh-",
	"ssh.",
	"review",
	"vimrc",
	"driver",
	"driven",
	"blog",
	"particle",
	"table",
	"intro",
	"importer",
	"thrift",
	"xmpp",
	"framework",
	"refresh",
	"react",
	"font",
	"librarie",
	"variou",
	"formatter",
	"analysi",
	"karma",
	"scroll",
	"tut_",
	"tut-",
	"tut.",
	"apple",
	"tag_",
	"tag-",
	"tag.",
	"tab_",
	"tab-",
	"tab.",
	"category",
	"ionic",
	"cache",
	"homebrew",
	"reverse",
	"english",
	"getting",
	"shipping",
	"clojure",
	"boot",
	"book",
	"branch",
	"combination",
	"combo",
}

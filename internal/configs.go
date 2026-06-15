package internal

// Configs is here

/*
app:
  name: "Chimera Agent"
  version: "1.0.0"
  #  date: "2025-01-31"
  environment: "development"
  debug: true

server:
  host: "localhost"
  port: 8787
  read_timeout: 15s
  write_timeout: 15s

runtimes:
  python: "./runtimes/python-3.14.2/.venv/bin/python"
  node_modules: "./runtimes/node-v25.4.0/node_modules"
  node_bin: "/usr/bin/node"

llm:
  provider: "ollama"
  host: "http://localhost:11434"
  model_main: "qwen3-vl:8b"
  formatter_model: "qwen3-vl:8b"
  temperature: 0.7
  formatter_temperature: 0.2
  format: "json"
  formatter_format: "json"
  max_conversation_messages: 25
#  context_window: 8192

# services:
#   python_analyzer:
#     url: "http://localhost:8000"
#     timeout: 30s
#
#   node_browser:
#     url: "http://localhost:3001"
#     timeout: 45s

# telegram:
#   enabled: true
#   token: "ENV_TELEGRAM_TOKEN"
#   allowed_users: [123456789]

paths:
  prompts_dir: "./configs/prompts"
  database: "./data/db/chimera.db"
  logsdir: "./data/logs"
*/

// 1. Define the Config struct
type Config struct {
    App struct {
        Name        string `yaml:"name"` 
        Version     string `yaml:"version"`
        Environment string `yaml:"environment"`
        Debug       bool   `yaml:"debug"`
    } `yaml:"app"`

    Server struct {
        Host         string `yaml:"host"`
        Port         int    `yaml:"port"`
        ReadTimeout  string `yaml:"read_timeout"`
        WriteTimeout string `yaml:"write_timeout"`
    } `yaml:"server"`

    Runtimes struct {
        Python      string `yaml:"python"`
        NodeModules string `yaml:"node_modules"`
        NodeBin     string `yaml:"node_bin"`
    } `yaml:"runtimes"`

    LLM struct {
        Provider                string  `yaml:"provider"`
        Host                    string  `yaml:"host"`
        ModelMain               string  `yaml:"model_main"`
        FormatterModel          string  `yaml:"formatter_model"`
        Temperature             float64 `yaml:"temperature"`
        FormatterTemperature    float64 `yaml:"formatter_temperature"`
        Format                  string  `yaml:"format"`
        FormatterFormat         string  `yaml:"formatter_format"`
        MaxConversationMessages int     `yaml:"max_conversation_messages"`
    } `yaml:"llm"`

    Paths struct {
        PromptsDir string `yaml:"prompts_dir"`
        Database   string `yaml:"database"`
        LogsDir    string `yaml:"logsdir"`
    } `yaml:"paths"`
}

// 2. Declare global variable
var AppConfig Config

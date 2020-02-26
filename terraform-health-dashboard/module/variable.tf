variable "env" {
  type = string
}

variable "lambdas_folder" {
  type = string
}

variable "memory_size" {
  type    = number
  default = 128
}

variable "timeout" {
  type    = number
  default = 3
}

variable "slack_warning" {
  type = string
  default = "xxxxx"
}
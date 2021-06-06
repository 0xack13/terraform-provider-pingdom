provider "pingdom" {}

resource "pingdom_check" "test" {
    name = "MyPDDemo"
    paused = false
    resolution = 1
    url = "google.com"
    port = 80
}
terraform {

  required_providers {

    todotest = {
      version = "0.0.1"
      source  = "terraform.local/lemurdaniel/todotest"
    }
  }

  backend "local" {

  }

}


provider "todotest" {
  some_string_config = "A string config for the provider"
}


data "todotest_item" "data_read_test_1" {
  id = 1
}


output "todo" {
  value = data.todotest_item.data_read_test_1
}

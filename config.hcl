io_mode = "Testing string"

service "tcp" "tcp/ip" {
    listen_addr = "10.10.10.10"

    process "command1" {
        command = ["date"]
    }
    process "command2" {
        command = ["date"]
    }
    
}
resource "null_resource" "res" {
  depends_on = [
	aws_dynamodb_table_item.init
  ]
  provisioner "local-exec" {
	when = create
	interpreter = ["/bin/bash","-c"]
	command = <<EOF

    declare -i timeout_max=${var.creation_delay}
    declare -i timeout=0
    while [[ $timeout -lt $timeout_max ]]; do
      sleep 1; timeout+=1 ; echo "*** wait  1 sek ($timeout) of $timeout_max"
    done

EOF
  }
    provisioner "local-exec" {
	when = destroy
	interpreter = ["/bin/bash","-c"]
	command = <<EOF
    declare -i timeout_max=${var.destroy_delay}
    declare -i timeout=0
    while [[ $timeout -lt $timeout_max ]]; do
      sleep 1; timeout+=1 ; echo "*** wait  1 sek ($timeout) of $timeout_max"
    done
    EOF
  }

}


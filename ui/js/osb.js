function executeRestOperation(actionOp) {
    if(actionOp.getAttribute("id") === "getBackups") {
        get_rest_ops("backups")
    }
    if(actionOp.getAttribute("id") === "getRestores") {
        get_rest_ops("restores")
    }
    if(actionOp.getAttribute("id") === "createBackup") {
        create_rest_ops("new_backup", "backups")
    }
    if(actionOp.getAttribute("id") === "createRestore") {
        create_rest_ops("new_restore", "restores")
    }
    if(actionOp.getAttribute("id") === "getBackup") {
        modify_backup_rest("backups")
    }
    if(actionOp.getAttribute("id") === "updateBackup") {
        create_rest_ops("update_backup", "backups")
    }
    if(actionOp.getAttribute("id") === "deleteBackup") {
        modify_backup_rest("delete")
    }
}

function get_actions() {
    var actions = []
    return new Promise(function(resolve, reject) {
        $.ajax({
        url: document.getElementById("get_actions").value,
        type: "GET",
        dataType: "json",
        error: function(xhr, status, exception){alert(xhr.responseText + xhr.status);},
        success: function (data) {
        paths = data.paths
        for(var path in paths){
            for ( var operation in paths[path]) {
                for (var operations in paths[path][operation]) {
                    var opId = paths[path][operation]["operationId"]
                    if(opId === "createServiceInstance") {
                        continue;
                    }
                    if(actions.length === 0) {
                        actions.push(opId)
                    }
                    var doesOpIdExist = false;
                    for(opIds in actions) {
                        if( opId == actions[opIds]) {
                            doesOpIdExist = true;
                        }
                    }
                    if(! doesOpIdExist) {
                        actions.push(opId)
                    }
                }
            }
        }
        resolve(actions)
        }
        });
    });
}


function displayActions() {
    $('#service_actions').children().empty()
    $('#service_actions').hide()
    //Parse the swagger and populate the actions list
    var instance_id = $("#get_instance_id").val();
    var url = $("#get_actions").val()
    if(instance_id === "") {
        alert("Service instance ID needs to be provided ")
        return;
    }
    if(url === "") {
        alert("url needs to be provided")
        return;
    }
    var actionsPromise =  get_actions()
    actionsPromise.then(function(actions){
        $('#service_actions').show()
        for(action in actions) {
            var button = document.createElement("button")
            var opName = actions[action]
            button.setAttribute("type", "button")
            button.setAttribute("id", opName)
            button.setAttribute("class", "btn btn-primary")
            button.setAttribute("onClick", "executeRestOperation(" + opName + ")")
            button.innerText = actions[action]
            $('#action_buttons').append(button)
            $('#action_buttons').append(document.createElement("br"))
            $('#action_buttons').append(document.createElement("br"))
        }
        $('#actions_header').text("Available actions for service: " + $("#get_instance_id").val());
    });

}

function clearActions(){
    window.location.href = "index.html"
}

$(function(){
	$(document).popover({
      selector: '[data-toggle=popover]',
      trigger: 'focus',
      html : true,
      content: function() {
        return $("#detail_div").html();
    }
  });
});

function createForm(table, rowNo, InputParam, cellNo) {
    var row = table.insertRow(rowNo)
    row.insertCell(cellNo).innerHTML = InputParam
    var input = document.createElement("input")
    input.setAttribute("id", InputParam)
    input.setAttribute("type", "text")
    input.setAttribute("size", "50")
    row.insertCell(cellNo + 1).appendChild(input)
}

function printDetails(object, id) {
    $('#detail_div').html("")
    var details_table = document.createElement("table")
    details_table.setAttribute("class", "table table-bordered")
    var params = Object.keys(object)
    for(var i=0; i<params.length; i++) {
        var param = params[i]
        var detail_row = details_table.insertRow(i)
        detail_row.insertCell(0).innerHTML = param
        detail_row.insertCell(1).innerHTML = object[param]
        if(i === params.length-1) {
            $('#detail_div').append(details_table)
            $(document.getElementById(id)).attr('data-content',$('#detail_div').html()).popover('show');
        }
    }
}

function getRestParams(object) {
    return new Promise(function(resolve, reject) {
        $.getJSON(document.getElementById("get_actions").value, function(data){
            resolve(Object.keys(data["definitions"][object]["properties"]))
        })
    })
}

function showDetails(link, object, id) {
    link.addEventListener('click', function () {
        printDetails(object, id)
    })
}

function get_rest_ops(action){
    $("#actions_result_div").html("");
    var table = document.createElement("table")
    table.setAttribute("id", action + "_table")
    table.setAttribute("class", "table table-bordered")
    var instance_name = document.getElementById("get_instance_id").value
    var base_url = document.getElementById("get_actions").value.split("/v2")[0]
    var url = base_url + "/v2/service_instances/" + instance_name + "/" + action

    $.ajax({
        url:url,
        type: "GET",
        dataType:"json",
        error: function(xhr, status, exception){alert(xhr.responseText + xhr.status);},
        success: function (data) {
            //Remove non null values from list
            data = data.filter(function(value) {
                return value != null
            })

            if (data.length === 0) {
                alert("There are no " + action + " for this instance")
                return;
            }
            for(var i =0; i<data.length; i++) {
                var data_link = document.createElement("a")
                var id = null
                var object = data[i]

                if(action === "backups") {
                    id = data[i].backup_id
                    data_link.setAttribute("id", id)
                    data_link.innerText = id
                }
                else if(action == "restores") {
                    id = data[i].restore_id
                    data_link.setAttribute("id", id)
                    data_link.innerText = id
                }
                data_link.setAttribute("data-toggle", "popover")
                data_link.setAttribute("href", "#")
                showDetails(data_link, object, id)
                table.insertRow(i).insertCell(0).appendChild(data_link)
                if(i === data.length -1) {
                    $('#actions_result_div').append(table)
                }
            }
        }
    })
}

function commit_creation(action, end_point) {
    var instance_name = document.getElementById("get_instance_id").value
    var base_url = document.getElementById("get_actions").value.split("/v2")[0]
    var url = base_url + '/v2/service_instances/' + instance_name + '/' + end_point
    var method = 'POST'
    if (action === "update_backup") {
        var backup_id = document.getElementById("Update_Backup_id:").value
        url = url + "/" + backup_id
        method = 'PUT'
    }
    getRestParams(action).then(function(data) {
        parameters = {}
        for(var i=0;i<data.length ;i++ ) {
            parameters[data[i]] = document.getElementById(data[i] + ":").value
            if(i===data.length-1 ) {
                $.ajax({
                    url: url,
                    type: method,
                    data: JSON.stringify(parameters),
                    dataType: "json",
                    contentType: "application/json",
                    error: function(xhr, status, exception){alert(xhr.responseText + xhr.status);},
                    success: function (data) {
                        get_rest_ops(end_point)
                    }
                });
            }
        }
    });
}

function create_rest_ops(action, end_point){
    $("#actions_result_div").html("");
    getRestParams(action).then(function(data) {
        var form = document.createElement("form")

        form.addEventListener("submit", function(event) {
            event.preventDefault()
            commit_creation(action, end_point)
        })
        var table = document.createElement("table")
        table.setAttribute("id", "take_" + action + "_table")
        var tableRow = 0
        for(var i=0; i<data.length; i++) {

            if(action === "update_backup" && i===0) {
                createForm(table, 0, "Update_Backup_id:" , 0)
                tableRow = tableRow + 1
            }
            createForm(table, tableRow, data[i] + ":", 0)
            tableRow = tableRow + 1

            if(i === data.length - 1) {
                form.appendChild(table)

                var submit = document.createElement("input")
                submit.setAttribute("type", "submit")
                submit.setAttribute("value", "Submit")

                form.appendChild(document.createElement("br"))
                form.appendChild(submit)

                $('#actions_result_div').append(form)
            }
        }
    })

}

function modify_backup(action) {
    var backup_id = document.getElementById("Backup_id:").value
    var instance_name = document.getElementById("get_instance_id").value
    var base_url = document.getElementById("get_actions").value.split("/v2")[0]
    var backup_url = base_url + "/v2/service_instances/" + instance_name + "/backups/" + backup_id
    var method = 'GET'
    var expectedType = 'json'
    if (action === 'delete') {
        method = 'DELETE'
        expectedType = 'text'
    }
    $.ajax({
        url: backup_url,
        type: method,
        dataType: expectedType,
        success: function(backup_data) {
            if(action === 'delete') {
                get_rest_ops('backups')
            }
            else {
                $("#actions_result_div").html("")
                var table = document.createElement("table")

                var backup_link = document.createElement("a")
                backup_link.setAttribute("id", backup_id)
                backup_link.setAttribute("data-toggle", "popover")
                backup_link.setAttribute("href", "#")
                backup_link.innerText = backup_id
                backup_link.addEventListener('click', (function(){
                    printDetails(backup_data, backup_id)
                }))

                table.insertRow(0).insertCell(0).appendChild(backup_link)
                $("#actions_result_div").append(table)
            }
        },
        error: function(xhr, status, exception){alert(xhr.responseText + xhr.status);},

    })
}

function modify_backup_rest(action) {
    $("#actions_result_div").html("");
    var form = document.createElement("form")
    form.addEventListener("submit", function(event) {
        event.preventDefault()
        modify_backup(action)
    })
    var table = document.createElement("table")
    createForm(table, 0, "Backup_id:" , 0)
    form.appendChild(table)
    var submit = document.createElement("input")
    submit.setAttribute("type", "submit")
    submit.setAttribute("value", "Submit")
    form.appendChild(document.createElement("br"))
    form.appendChild(submit)
    $('#actions_result_div').append(form)
}

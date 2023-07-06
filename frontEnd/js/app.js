getData();
function removeFromDb(id) {
    // if (confirm("Вы уверены, что хотите удалить этого студента?")) {
        confirmDelete(id);
        location.href = 'main.html';
    // }
}

const confirmDelete = async (id) => {

        try {
            const response = await fetch("http://localhost:8888/delete/" + id, {
                method: 'DELETE'
            })
            const responseJson = await response.json();
        }
        catch (e) {
            console.log(e)
        }

}

async function getData() {
    const response = await fetch('http://localhost:8888/');
    console.log(response);
    const data = await response.json();
    console.log(data);
    const cat = document.getElementById("data");
    data.forEach(element => {
        cat.innerHTML += `
        <div class="bg-body-tertiary p-5 rounded mt-3">
          <h3 >${element.Activity}</h3>
          <p class="lead modal">${element.Category}</p> 
          <p class="lead">Участники: ${element.Participants}</p> 
          <p><a href="${element.Link}" target="_blank" class="link-primary link-offset-2 link-underline-opacity-25 link-underline-opacity-100-hover">${element.Link}</a></p>  
          <a onclick="join(${element.ID})" class="btn btn-primary" href="#" role="button">Участвовать</a>
          <a onclick="openModal()" class="btn btn-success img" role="button">Изменить</a>
          <a onclick="removeFromDb(${element.ID})"  class="btn btn-danger" role="button">Удалить</a>
        </div>   
        <div id="myModal" class="modal bg-body-tertiary p-5 rounded mt-3 modal-sm">
        <div class="modal-content">
            <span class="close" onclick="closeModal()">&times;</span>
            <input name="text" for="text" placeholder="text" type="text" id="text">
            <button onclick="saveChanges(${element.ID})">Save</button>
            <button onclick="cancelChanges()">Cancel</button>
        </div>
    </div>
`
    });
}

function openModal() {
    var modal = document.getElementById("myModal");
    modal.style.display = "block";
}

function closeModal() {
    var modal = document.getElementById("myModal");
    modal.style.display = "none";
}
function saveChanges(id) {
    var modal = document.getElementById("myModal");
    var inputElement = document.getElementById("text");
    // Save the changes made in the input field
    var text = inputElement.value;

    // Convert the form data to a JSON object
    const newText = {text};

    console.log('newText ==> ',newText)
    // Send the form data to the server using Fetch API
    fetch('http://localhost:8888/editText/' + id, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(newText)
    })
        .then(result => {
            location.href = 'http://localhost:63342/activity/frontEnd/main.html';
            console.log(result);
            // Handle the response from the server
        })
        .catch(error => {
            console.error(error);
            // Handle any error that occurred during the request
        });
    // Update the text in your HTML page with the new value
    // For example, assuming you have a <div> element with id "myText":
    var textElement = document.getElementById("myText");

    // Hide the modal
    modal.style.display = "none";
}
function cancelChanges() {
    var modal = document.getElementById("myModal");
    modal.style.display = "none";
}

function join(id){
 confirmJoin(id)
    location.href = 'main.html';
}

const confirmJoin = async (id) => {
    try {
        const response = await fetch("http://localhost:8888/join/" + id , {
            method: 'POST'
        })
        const responseJson = await response.json();
        console.log(responseJson)
    }
    catch (e) {
        console.log(e)
    }
}

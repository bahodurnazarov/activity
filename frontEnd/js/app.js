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

    var temp = "";


    const cat = document.getElementById("data");
    data.forEach(element => {
        cat.innerHTML += `
        <div class="bg-body-tertiary p-5 rounded mt-3" >
          <h3>${element.Activity}</h3>
          <p class="lead">${element.Category}</p> 
          <p class="lead">Участники: ${element.Participants}</p> 
          <p><a href="${element.Link}" target="_blank" class="link-primary link-offset-2 link-underline-opacity-25 link-underline-opacity-100-hover">${element.Link}</a></p>  
          <a class="btn btn-lg btn-primary" href="/docs/5.3/components/navbar/" role="button">Участвовать</a>
          <a onclick="removeFromDb(${element.ID})" class="btn btn-lg btn-danger" role="button">Удалить</a>
        </div>`
    });

}
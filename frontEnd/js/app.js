getData();

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
          <p><a href="${element.Link}" class="link-primary link-offset-2 link-underline-opacity-25 link-underline-opacity-100-hover">${element.Link}</a></p>  
          <a class="btn btn-lg btn-primary" href="/docs/5.3/components/navbar/" role="button">Участвовать</a>
        </div>`
    });

}
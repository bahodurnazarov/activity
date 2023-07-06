document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault();

    // Create a new FormData object
    const formData = new FormData(this);
    let token
    // Convert the form data to a JSON object

    function returnStr() {
        console.log(token);
        return token
    }
    const data = {};
    formData.forEach(function(value, key) {
        data[key] = value;
        console.log(value)
    });
    // Send the form data to the server using Fetch API
    fetch('http://localhost:8888/login', {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(result => {
            //location.href = 'http://localhost:63342/activity/frontEnd/main.html';
            token = result.token
            console.log(result);
            if (result.errorCode == "200") {
                //location.href = 'http://localhost:63342/activity/frontEnd/main.html';
                returnStr()
            }
            // Handle the response from the server
        })
        .catch(error => {
            console.error(error);
            // Handle any error that occurred during the request
        });
});

export default register
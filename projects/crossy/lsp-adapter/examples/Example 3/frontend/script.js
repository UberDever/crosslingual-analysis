function updateCounter() {
    // basically http://localhost:3333/item.POST
    fetch('http://localhost:3333/item', {
        method: "POST",
    })

    // basically http://localhost:3333/item.GET
    fetch('http://localhost:3333/item')
        .then(response => {
            console.log(response)
            return response.json()
        })
        .then(data => {
            document.getElementById('data').innerText = "The current value of counter is: " + data.count;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}
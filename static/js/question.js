async function sendQuestion() {
    const title = document.getElementById('title').value;
    const content = document.getElementById('content').value;
    let permission = document.getElementById('permission').value;
    permission = parseInt(permission);

    const response = await fetch('/api/question', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({'title': title, 'content': content, 'permission': permission}),
    })

    const data = await response.json();
    console.log(data);
}

async function getQuestionTitleAndContent() {
    let path = window.location.pathname;
    let segments = path.split('/');
    const id = segments[2];
    if (!/^\d+$/.test(id)) {
        window.location.href = '/404';
        return;
    }
    const url = 'api/question/' + id;

    const response = await fetch(url, {method : 'GET'})
    const data = await response.json();
    if(!response.ok) {
        alert(data['error']);
    } else {
        document.getElementById('title').value = data['title'];
        document.getElementById('content').value = data['content'];
        document.getElementById('permission').value = data['permission'];
    }
}

async function modifyQuestion() {
    const title = document.getElementById('title').value;
    const content = document.getElementById('content').value;
    let permission = document.getElementById('permission').value;
    permission = parseInt(permission);
    
}
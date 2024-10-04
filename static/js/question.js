function getQuestionUrl() {
    const path = window.location.pathname;
    const segments = path.split('/');
    const id = segments[2];
    if (!/^\d+$/.test(id)) {
        return '';
    }
    return window.location.origin + '/api/question/' + id;
}

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

function getQuestionTitleAndContent() {
    let path = window.location.pathname;
    let segments = path.split('/');
    const id = segments[2];
    if (!/^\d+$/.test(id)) {
        window.location.href = '/404';
        return;
    }
    const url = window.location.origin + '/api/question/' + id;

    fetch(url, {method : 'GET'})
        .then((response) => {
            return response.json().then((data) => {
                if (!response.ok) {
                    alert(data['error']);
                    window.history.back()
                    return Promise.reject('Failed to fetch');
                }
            return data;
            });
        })
        .then((data) => {
            document.getElementById('title').value = data['title'];
            document.getElementById('content').value = data['content'];
            document.getElementById('permission').value = data['permission'];
        })
        // .catch((error) => console.error('Error:', error));
}

async function modifyQuestion() {
    const url = getQuestionUrl();
    if (!url) {
        window.location.href = '/404';
        return;
    }

    const title = document.getElementById('title').value;
    const content = document.getElementById('content').value;
    let permission = document.getElementById('permission').value;
    permission = parseInt(permission);

    const response = await fetch(url, {
        method: 'PUT',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({"id": parseInt(id), "title": title, 'content': content, 'permission': permission}),
    })
    const data = await response.json();
    if(!response.ok) {
        alert(data['error']);
    } else {
        console.log(data);
    }
}

async function deleteQuestion() {
    const url = getQuestionUrl();
    if (!url) {
        window.location.href = '/404';
        return;
    }

    const response = await fetch(url, {
        method: 'DELETE',
    })
    const data = await response.json();
    console.log(data);
}
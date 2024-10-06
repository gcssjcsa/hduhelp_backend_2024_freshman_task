function getQuestionUrl() {
    const path = window.location.pathname;
    const segments = path.split('/');
    const qid = segments[2];
    if (!/^\d+$/.test(qid)) {
        return '';
    }
    return window.location.origin + '/api/question/' + qid;
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
    });

    const data = await response.json();
    console.log(data);
}

async function getQuestion() {
    const url = getQuestionUrl();
    if (!url) {
        window.location.href = '/404';
        return;
    }

    try {
        const response = await fetch(url);
        const data = await response.json();
        if (!response.ok) {
            alert(data['error']);
            window.history.back();
            return Promise.reject('Failed to fetch');
        }
        return data;
    } catch (error) {
        console.error('Error fetching question:', error);
    }
}

async function modifyQuestion() {
    const path = window.location.pathname;
    const segments = path.split('/');
    const qid = segments[2];
    if (!/^\d+$/.test(qid)) {
        return '';
    }
    const url = window.location.origin + '/api/question/' + qid;

    const title = document.getElementById('title').value;
    const content = document.getElementById('content').value;
    let permission = document.getElementById('permission').value;
    permission = parseInt(permission);

    const response = await fetch(url, {
        method: 'PUT',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({"id": parseInt(qid), "title": title, "content": content, "permission": permission}),
    });
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
    });
    const data = await response.json();
    console.log(data);
}

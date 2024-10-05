function getQuestionId() {
    const path = window.location.pathname;
    console.log(path);
    const segments = path.split('/');
    const qid = segments[2];
    if (!/^\d+$/.test(qid)) {
        return null;
    }
    return qid;
}

function getAnswerId() {
    const path = window.location.pathname;
    const segments = path.split('/');
    const aid = segments[4];
    if (!/^\d+$/.test(aid)) {
        return null;
    }
    return aid;
}

function getIds() {
    const qid = getQuestionId();
    if (!qid) {
        window.location.href = '/404';
        return null;
    }
    const aid = getAnswerId();
    if (!aid) {
        window.location.href = '/404';
        return null;
    }
    return { qid, aid };
}

async function sendAnswer() {
    const qid = getQuestionId();
    if (!qid) {
        window.location.href = '/404';
        return;
    }
    const url = `${window.location.origin}/api/question/${qid}/answer`;

    const content = document.getElementById('content').value;

    const response = await fetch(url, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({'content': content}),
    })

    const data = await response.json();
    console.log(data);
}

function getAnswerContent() {
    const ids = getIds();
    if (ids) {
        const url = `${window.location.origin}/api/question/${ids.qid}/answer/${ids.aid}`;

        fetch(url, {method: 'GET'})
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
                document.getElementById('content').value = data['content'];
            })
        // .catch((error) => console.error('Error:', error));
    }
}

async function modifyAnswer() {
    const ids = getIds();
    if (ids) {
        const url = `${window.location.origin}/api/question/${ids.qid}/answer/${ids.aid}`;
        const content = document.getElementById('content').value;

        const response = await fetch(url, {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({"id": parseInt(ids.aid), 'content': content}),
        });
        const data = await response.json();
        if (!response.ok) {
            alert(data['error']);
        } else {
            console.log(data);
        }
    }
}

async function deleteAnswer() {
    const ids = getIds();
    if (ids) {
        const url = `${window.location.origin}/api/question/${ids.qid}/answer/${ids.aid}`;

        const response = await fetch(url, {
            method: 'DELETE',
        })
        const data = await response.json();
        console.log(data);
    }
}
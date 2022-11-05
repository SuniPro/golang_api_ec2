function dailyCheck(){
    return fetch('/api/admin/daily_check')
        .then(response => {
            return response.json();
        })
        .then(dailyCheck => {
            return dailyCheck;
        })
        .catch(error => console.log(error))
}

function reward(){
    return fetch('/api/admin/compensation')
        .then(response => {
            return response.json();
        })
        .then(reward => {
            return reward;
        })
        .catch(error => console.log(error))
}

function checkCounting(){
    return fetch('api/admin/counting_check')
        .then(response => {
            return response.json();
        })
        .then(counting => {
            return counting;
        })
        .catch(error => console.log(error))
}

export {
    dailyCheck,
    reward
}
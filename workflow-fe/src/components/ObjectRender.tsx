function ObjectRender() { 
    const task = { 
        title: "เบิกงบ", 
        amount: 20, 
    }; 
    return ( 
    <div> 
        <h1>Task: {task.title}</h1> 
        <p>${task.amount}</p> 
        </div> 
        ); 
    } 
    export default ObjectRender;
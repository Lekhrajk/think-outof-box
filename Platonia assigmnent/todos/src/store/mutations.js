 export const SET_TODOS = (state,todos) =>{
     state.Todos = todos
 }

export const SET_TODO = (state,todo) =>{
    state.Todo = todo;
}

export const ADD_TO_TODO = (state,{id,userId,title,completed}) =>{
    state.Todos.push({
        id,
        userId,
        title,
        completed
    })
    return state.Todos.reverse();
}


export const REMOVE_TODO =(state,id) => {
    state.Todos = state.Todos.filter(item =>{
        return item.id!==id;
    })
}


import axios from 'axios';

export const getTodos =  ({ commit }) => {
    axios.get("https://jsonplaceholder.typicode.com/todos/")
        .then(response => {
           commit('SET_TODOS', response.data);
        }) 
} 

export const getTodo =  ({ commit } ,todoId) => {
    axios.get(`https://jsonplaceholder.typicode.com/todos/${todoId}`)
        .then(response => {
           commit('SET_TODO', response.data);
        }) 
} 


export const addTodo =  ({ commit} ,{id,userId,title,completed}) => {
     commit('ADD_TO_TODO',{id,userId,title,completed});
} 

export const removeTodo = ({commit} , id) =>{
   commit( 'REMOVE_TODO',id);
   axios.delete(`https://jsonplaceholder.typicode.com/todos/${id}`);
}


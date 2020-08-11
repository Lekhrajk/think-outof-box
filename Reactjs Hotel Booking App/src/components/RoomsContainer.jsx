import React from 'react'
import RoomsFilter from './RoomsFilter';
import RoomsList from './RoomsList';
import { withRoomConsumer } from '../context';
import Loading from './Loading';

function RoomsContainer({context}){
    const{loading,sortedRooms,rooms} = context;
    if(loading)
    {
        return <Loading/>;
    }
    return(
        <>
            <RoomsFilter rooms={rooms}/>
            <RoomsList rooms = {sortedRooms}/>
        </>
    );

}
export default withRoomConsumer(RoomsContainer)

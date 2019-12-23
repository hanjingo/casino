import { Table } from "./table";
import { Player } from "../../../lib/player";

/**
 * 房间
 */
export class Room {
    //房间id
    private _id: number;

    //桌子
    private _table: Map<number, Table>;

    //玩家集合
    private _players: Map<number, RoomPlayer>;
}

/**
 * 房间里的玩家
 */
export class RoomPlayer extends Player {

}
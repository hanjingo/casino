import { Seat } from "./seat";
import { Round } from "./round";

/**
 * 桌子
 */
export class Table {
    //座位集合
    private _seats: Map<number, Seat>;
    //回合集合
    private _rounds:Map<number, Round>;
}
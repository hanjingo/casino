/**
 * 玩家
 */
export class Player {
    //玩家id
    private id:number;
    //玩家名字
    private name:string;
    //玩家头像
    private headImgUrl:string;
    //性别
    private sex:Sex;
    //机密
    private secret:PlayerSecret;
}

//玩家机密信息
class PlayerSecret {
    //密码
    private passWord:string;
    //token
    private token:string;
    //手机
    private phone:string;
    //email
    private email:string;
}

cc.Class({
    extends: cc.Component,

    properties: {
        //筹码对象
        obj: {
            default: null,
            type: cc.Button,
        },

        //筹码数值
        num: {
            default: null,
            type: cc.Label,
        },

        //筹码背景图片
        bg: {
            default: null,
            type: cc.Sprite,
        },

        //筹码状态变量
        doMove: { //移动
            default: false,
            type: cc.Boolean,
        },
        speed: { //移动速度
            default: 0.0,
            type: cc.Float,
        },
        startMoveTime: { //开始时间
            default: null,
            type: Number,
        },
        endMoveTime: { //结束时间
            default: null,
            type: Number,
        },
        endPoint: { //终点
            default: null,
            type: cc.Vec3,
        }
    },

    // LIFE-CYCLE CALLBACKS:

    // onLoad () {},

    start () {

    },

    //设置筹码
    setNum : function(n) {
        this.num.string = n;
    },

    //更新筹码
    doUpdate : function() {
        this.updateMove();
    },

    //更新移动
    updateMove: function() {

    },

    //移动到指定位置 (终点, 时长)
    moveTo: function(endPoint, durs) {
        this.endPoint = endPoint;
        this.startMoveTime = 
    },

    // update (dt) {},
});

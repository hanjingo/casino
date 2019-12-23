declare const enum CmpResult {
    Bigger = 1,
    Equal = 2,
    Smaller = 3,
}

interface PokerI {
    GetColor();
    GetPoint();
    CmpColor(pk:PokerI):CmpResult;
    CmpPoint(pk:PokerI):CmpResult;
}

export class Poker {
    //花色
    private color:number;

    //点数
    private point:number;

    constructor(color:number, point:number) {
        this.color = color;
        this.point = point;
    }

    //获得牌的花色
    public GetColor():number {
        return this.color;
    }

    //获得牌的点数
    public GetPoint():number {
        return this.point;
    }

    //比较花色
    public CmpColor(pk:Poker):CmpResult {
        return this.GetColor() > pk.GetColor() ? CmpResult.Bigger : 
            this.GetColor() == pk.GetColor() ? CmpResult.Equal : CmpResult.Smaller;
    }

    //比较点数
    public CmpPoint(pk:Poker):CmpResult {
        return this.GetPoint() > pk.GetPoint() ? CmpResult.Bigger : 
            this.GetPoint() == pk.GetPoint() ? CmpResult.Equal : CmpResult.Smaller;
    }
}

export class PokerSuit {
    //扑克集合
    private pokers:Array<PokerI>;

    constructor() {
        this.pokers = new Array<PokerI>();
    }

    //按点数从大到小排序
    public SortWithPointByDesc(pks:Array<PokerI>):void {
        pks.sort();//todo
    }

    //按点数从小到大排序
    public SortWithPointByAesc(pks:Array<PokerI>):void {
        pks.sort();//todo
    }

    //按颜色从大到小排序
    public SortWithColorByDesc(pks:Array<PokerI>):void {
        pks.sort();//todo
    }

    //按颜色从小到大排序
    public SortWithColorByAesc(pks:Array<PokerI>):void {
        pks.sort();//todo
    }

    //先按点数排序再把相同的点数按照颜色从大到小排序
    public SortWithPointAndColorByDesc(pks:Array<PokerI>):void {
        pks.sort();//todo
    }

    //先按点数排序再把相同点数按照颜色从小到大排序
    public SortWithPointAandColorByAesc(pks:Array<PokerI>):void {
        pks.sort();//todo
    }

    //找到点数最小的牌
    public GetMinCardWithPoint(pks:Array<PokerI>):Array<PokerI> {
        let back = new Array<PokerI>();
        //todo
        return back;
    }

    //找到点数最大的牌
    public GetMaxCardWithPoing(pks:Array<PokerI>):Array<PokerI> {
        let back = new Array<PokerI>();
        //todo
        return back;
    }
}
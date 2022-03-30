/**
 @constructor
 @abstract
 */
class Exp  {
    constructor() {
        if (this.constructor == Exp) {
        throw new Error("Abstract classes can't be instantiated.");
        }
    }
    eval(){ console.log("Exp eval")}
};

/**
 @constructor
 @abstract
 */
class LitExp extends Exp {
    constructor(val) {
        if (this.constructor == LitExp) {
           throw new Error("Abstract classes can't be instantiated."); 
        }
        super()
        this.val = val
        this.type = null
    }


    eval() {
        assert(typeof this.val == this.type && this.type != null)
        return this.val
    }
}

class StringLit extends LitExp { type = String }
class IntLit extends LitExp { type = Number}
class BoolLit extends LitExp { type = Boolean }
class DoubleLit extends LitExp { type = Number }
class FloatLit extends LitExp { type = Number }

/**
 @constructor
 @abstract
 */
class BinOp extends Exp {
    constructor(leftExp, rightExp) {
        if (this.constructor == BinOp) {
            throw new Error("Abstract classes can't be instantiated")
        }
        super()
        this.leftExp = leftExp
        this.rightExp = rightExp
    }

    logic(left, right) { 
        throw new Error(`BinOp logic not implemented for ${this.constructor.name}`)
    }

    eval() { return this.logic(this.leftExp.eval(), this.rightExp.eval()) }
}

class PlusBinOp extends BinOp { logic(left, right) { left + right }}
class MinusBinOp extends BinOp { logic(left, right) { left - right }}
class DivideBinOp extends BinOp { logic(left, right) { left / right }}
class MultiplyBinOp extends BinOp { logic(left, right) { left * right }}

class TagExp extends Exp {
    constructor(exp, tag) {
        super()
        this.exp = exp 
        this.tag = tag
    }
}

class PackageExp extends Exp {
    constructor(exp, packTag) {
        super()
        this.exp = exp
        this.packTag = packTag
    }
}
module.exports = {
    Literals : {
        StringLit: this.StringLit,
        IntLit: this.IntLit,
        BoolLit: this.BoolLit,
        DoubleLit: this.DoubleLit,
        FloatLit: this.FloatLit
    },
    BinOpExp : {
        PlusBinOp: this.PlusBinOp,
        MinusBinOp: this.MinusBinOp,
        MultiBinOp: this.MultiBinOp,
        DivideBinOp: this.DivideBinOp,
        
    },
    TagExp: this.TagExp,
    PackageExp: this.PackageExp
}
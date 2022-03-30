const EvalHandler = require("./EvalHandler")
const {Literals, BinOpExp} = require("../model/Exp")
const { netLog } = require("electron")
class ParseHandler {
    static pack = "STANDARD"
    static parse(input, parsed) {
        function inParse(e, acc, t) {
            var acc = new Array()
            function inLogic(e, acc) {
                if (e == null) {
                    return acc
                }
                if (/^ +$/.test(e.x)) {
                    let t = acc.push(inParse(e.next))
                    return {w: e.x, l: t} 
                } else {
                    let n = inLogic(e.next)
                    return {w: n.w + e.x,l: n.l}
                }
            }
            if (e == null) {
                return null
            } else {
                let n = inLogic(e, acc)
                acc.push(n.w)
                return n.l.push(acc)
            }
        }


        let exp = inParse(input, new Array())
        parsed(exp)
        return EvalHandler.eval(ParseHandler.construcExp(exp), ParseHandler.pack)
    }

    

    static construcExp(e) {
        return e 

    }
}

module.exports = ParseHandler
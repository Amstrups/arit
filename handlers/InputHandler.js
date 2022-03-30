const Cons = require('../model/Cons')

class InputHandler {
    constructor() { this.stack = null}
    
    clear() {
        this.stack = null
    }

    async save() {
        console.log("Saving input")
    }

    push(char) {
        this.stack = new Cons(char, this.stack)
    }

    pop() {
        if (this.stack != null) {
            this.stack = this.stack.next
        }
    }
    // Testing with regex if true
    regTesting = true

    eval(input, shouldEval, shouldAppend, shouldRemove) {
        if (input.type == "keyDown" && this.regTesting) {
            switch(true) {
                case /(^[a-z]+$|^[A-Z]+$)/.test(input.key):
                case /^[0-9]+$/.test(input.key):
                case /^\++$/.test(input.key):
                case /^ +$/.test(input.key):
                    this.push(input.key)
                    shouldAppend(input.key)
                    break
                case /(e|E)nter/.test(input.code):
                    shouldEval(this.stack)
                    this.save()
                    this.clear()
                    break
                case /(b|B)ackspace/.test(input.code):
                    this.pop()
                    shouldRemove()
                    break
                default:
                    console.log("Found other:")
                    console.log(input)
                    break
            }
        } else {
            if (input.location < 1 && input.type == "keyDown") {
                console.log(input)
            }
        }
    }
}

module.exports = InputHandler
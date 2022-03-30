class Cons {
    constructor(x_val = null, next_val = null) {
        this.x = x_val
        this.next = next_val
    }

    append(x_val) {
        return new Cons(x_val, this)
    }
}

Cons.prototype.toString = function() {
     return `\n x: ${this.x}, next: ${this.next}`
}

module.exports = Cons
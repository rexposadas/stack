/**
 * @param {...(null|boolean|number|string|Array|Object)} args
 * @return {number}
 */
var argumentsLength = function(...args) {

    if (args === undefined) {
        return 0
    }
    return args.length

};

/**
 * argumentsLength(1, 2, 3); // 3
 */


// This one was pretty easy. took less than 5 minutes. had to look up how to guard from undefined..

function testResolve() {
    const a = new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve("a")
        }, 2000)
    })

    /*a.then((result) => {
        console.log(result);
        return result + 'b'
    })*/
    a.then((result) => {
        console.log(result);
        return Promise.resolve(result + 'b')
    })
    a.then((result) => {
        console.log(result)
    })
}

//testResolve();

function helper1() {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve('helper1')
        }, 2000)
    })
}

function helper2() {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve('helper2')
        }, 3000)
    })
}

function asyncFunction(cb) {
    return () => {
        const generatorObject = cb();

        function handler(result) {
            console.log('runing handler')
            if (result.done) {
                return Promise.resolve(result.value)
                //return result.value
            }

            //return Promise.resolve(result.value)
            return Promise.resolve(result.value).then(res => {
                console.log('insidw resolve');
                return handler(generatorObject.next(res))
            })
        }

        return handler(generatorObject.next())
    }
}

function *testAsyncGen1() {
    try {
        //return 0
        yield helper1()
        yield 0
        return helper2()
        //return 3;
    } catch(err) {
        console.log(err)
    }
}

const final = asyncFunction(testAsyncGen1);
final().then( (res) => console.log(res));
console.log('resolve returned');

function doT() {
    return Promise.resolve('').then((res) => console.log('dfirstonme')).then((res) => console.log('before second'))
}

doT().then((res) => console.log('second'))
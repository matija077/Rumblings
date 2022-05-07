const sendRequestButton = document.getElementById('send-request');
sendRequestButton.addEventListener('click', sendRequest);



/*function testResolve() {
    const a = new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve("a")
        }, 2000)
    })

    /*a.then((result) => {
        console.log(result);
        return result + 'b'
    })*/ /*
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
            if (result.done) {
                return Promise.resolve(result.value)
                //return result.value
            }

            //return Promise.resolve(result.value)
            return Promise.resolve(result.value).then(res => {
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

const { decorateEventHandler, onDestroy } = createDecorateEventHandler();
const inputText = document.getElementById('inputText');
inputText.addEventListener("input", decorateEventHandler("inputText", handleInput, [], {passEvent: true, throttle: {

} }));
function handleInput(value) {
    console.log(value);
}

onDestroy();

function createDecorateEventHandler() {
    var initialized, throttlingParams, throttlingParams;

    function decorateEventHandler(type, cb, args = [], options = {}) {
        if (!initialized) {
            onCreate();
        }

        if (!cb || !type) {
            return;
        }

        console.log('tu sam');
    
        const needsThrottling = ({ interval = 1000 }) => {
            var throttlingInProgress = true;

            let timeoutId = throttlingParams.get(type);
            if (!timeoutId) {
                throttlingParams.set(
                    setTimeout(
                        () => {
                            throttlingParams.delete(timeoutId);
                        }, 
                        interval
                    )
                )
            }
             
            return throttlingInProgress;
        }
        const needsDebounceing = () => {
    
        }
    
        let isDisabled = false;
    
        if (options.throttle) {
            isDisabled = needsThrottling(options.throttle)
        } else if (options.debounce) {
            isDisabled = needsDebounceing(options.debounce);
        }
    
        return (event) => {
            options.passEvent && (cb = cb.bind(null, event));
            console.log(options);

            !isDisabled && cb(...args);
        }
    }

    function onDestroy() {
        initialized = false;
        reset();
    }

    function onCreate() {
        initialized = true;
        initialize();
    }

    function isInitiliazed() {
        return initialized;
    }

    function reset() {
        for (p of throttlingParams) {
            clearTimeout(p)
        }
        for (p of debounceParams) {
            //clearTimeout(p)
        }
        initialize()
    }

    function initialize() {
        throttlingParams = new Map();
        debounceParams = new Map();
    }

    return {
        decorateEventHandler,
        onDestroy,
        onCreate,
        isInitiliazed
    }

}

const button = document.getElementById('testClick');
//button.onclick = () => console.log('from propert');
button.addEventListener('click', () => console.log('from event listener'))

*/

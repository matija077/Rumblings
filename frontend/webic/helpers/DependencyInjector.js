

export class DependencyInjector {
    constructor() {
        this.services = {};
    }

    register(name, service) {
        if (this.services[name]) {
            return;
        }
        
        this.services[name] = service;
    }

    resolve(dependencies, func) {
        return (...args) => {
            const resolved = dependencies.map(dependency => this.services[dependency]);

            return func(...resolved, ...args);
        }
    }
}

class AbstractHttp {
  constructor() {
    if (this.constructor === "AbstractHttp") {
      throw new Error("should be abstract");
    }
  }

  getData() {
    throw new Error("must be implemented");
  }
}

class Http extends AbstractHttp {
  constructor() {
    super();
  }

  async fetchHelper(url, options) {
    let data;
    let error;

    try {
      data = await this.fetchWithJson(url, options)
    } catch (err) {
      error = err;
      console.error(err)
    }

    return { data, error };
  }

  async fetchWithJson(url, options = {method: 'GET'}) {
    const response = await fetch(url, options);
    return response.json();
  }

  getData(url) {
    return this.fetchHelper(url)
  }
}

export default Http;

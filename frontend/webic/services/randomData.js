class RandomData {
  constructor(http) {
    this.http = http;
  }

  url = "https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb";

  async getRandomData() {
    const { data, error } = await this.http.getData(this.url);
    return { data: data?.results[0] ?? [], error };
  }
}

export default RandomData;

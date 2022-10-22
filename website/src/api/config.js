const dev_config = {
  // The URL of the server to use.
  SERVER_HOST: "http://localhost:8080/",
  // SERVER_HOST: "http://tirelease.pingcap.net/",
};

export function url(url) {
  return `${dev_config.SERVER_HOST}${url}`;
}

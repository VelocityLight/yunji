const config = {
  // The URL of the server to use.
  // SERVER_HOST: "/",
  SERVER_HOST: "http://127.0.0.1:8080/",
};

export function url(url) {
  return `${config.SERVER_HOST}api/v1/${url}`;
}

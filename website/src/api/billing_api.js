import { url } from "./config"

export async function fetchBillingByTagAndService({ tags = [], service }) {
  var tagQuery = ""
  if (tags.length > 0) {
    tagQuery = `
    ${tags.map((item) => {
      return "tags=" + item
    }).join("&")
      }`
  }

  var serviceQuery = ""
  if (service) {
    serviceQuery = service
  }

  var queryString = ""
  if (tagQuery.length > 0 || serviceQuery.length > 0) {
    queryString = [tagQuery, serviceQuery].join("&")
  }

  console.log(
    "fetchBiling",
    url(`bills/trend?${queryString}`)
  );

  return fetch(url(`bills/trend?${queryString}`))
    .then(async (res) => {
      const data = await res.json();
      return data.data;
    })
  // .catch((e) => {
  //   return e
  // });
}

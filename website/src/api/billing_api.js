import { url } from "./config"

export async function fetchBillingByTagAndService({ tags = [], service }) {
  var tagQuery = ""
  if (tags.length > 0) {
    tagQuery = `tags=`.concat(tags.join(","))
  }

  var serviceQuery = ""
  if (service && service.length > 0) {
    serviceQuery = `service=`.concat(service.join(","))
  }

  var queryString = tagQuery
  if (serviceQuery.length > 0) {
    queryString = (queryString.length > 0 ? (queryString + "&") : "") + serviceQuery
  }

  console.log(
    "fetchBilling",
    url(`bills/trend?${queryString}`)
  );

  return fetch(url(`bills/trend?${queryString}`))
    .then(async (res) => {
      const data = await res.json();
      return data;
    })
  // .catch((e) => {
  //   return e
  // });
}

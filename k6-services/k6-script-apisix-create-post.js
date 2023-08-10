// Auto-generated by the postman-to-k6 converter

import "./libs/shim/core.js";
import "./libs/shim/urijs.js";
import { group } from "k6";

export const options = {
  maxRedirects: 4,
  thresholds: {
    http_req_failed: ["rate<0.02"], // http errors should be less than 2%
    http_req_duration: ["p(95)<2000"], // 95% requests should be below 2s
  },
  scenarios: {
    constant_request_rate: {
      executor: "constant-arrival-rate",
      rate: 1000,
      timeUnit: "1s", // 10 iterations per second, i.e. 10 RPS
      duration: "1m",
      preAllocatedVUs: 100, // how large the initial pool of VUs would be
      maxVUs: 1000, // if the preAllocatedVUs are not enough, we can initialize more
    },
  },
};

const Request = Symbol.for("request");
postman[Symbol.for("initial")]({
  options,
  environment: {
    "base_url.kong": "http://127.0.0.1:8000",
    "base_url.post-be": "http://post-backend:4000",
    "base_url.comment-be": "http://comment-backend:4001",
    "base_url.apisix": "http://10.211.55.10:9080",
  },
});

export default function () {
  group("APISIX BACKEND", function () {
    group("POST", function () {
      postman[Request]({
        name: "Create Post",
        id: "bc95ad0e-df0c-4130-b1bd-bf746c7d0343",
        method: "POST",
        address: "{{base_url.apisix}}/post",
        data: "{{backend.create.post.request-body}}",
        pre() {
          const title = pm.variables.replaceIn("{{$randomUserName}}");
          const body = pm.variables.replaceIn("{{$randomLoremText}}");

          const requestJSON = {
            title: `POST ${title}`,
            body: body,
          };

          pm.environment.set(
            "backend.create.post.request-body",
            JSON.stringify(requestJSON)
          );
        },
        post(response) {
          pm.test("Status code is 200", function () {
            pm.response.to.have.status(200);
          });
        },
      });
    });
  });
}
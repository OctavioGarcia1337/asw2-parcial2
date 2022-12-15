import http from 'k6/http';
import { check, sleep } from "k6";

const isNumeric = (value) => /^\d+$/.test(value);

const default_vus = 300;

const target_vus_env = `${__ENV.TARGET_VUS}`;
const target_vus = isNumeric(target_vus_env) ? Number(target_vus_env) : default_vus;

export let options = {
  stages: [
      // Ramp-up from 1 to TARGET_VUS virtual users (VUs) in 5s
      { duration: "5s", target: target_vus },

      // Stay at rest on TARGET_VUS VUs for 10s
      { duration: "8s", target: target_vus },

      // Ramp-down from TARGET_VUS to 0 VUs for 5s
      { duration: "5s", target: 0 }
  ]
};

export default function () {
  const response = http.get("http://host.docker.internal/search=*_*", {headers: {Accepts: "application/json"}});
  check(response, { "status is 200": (r) => r.status === 200 });
  const response2 = http.get("http://host.docker.internal:8090/1", {headers: {Accepts: "application/json"}});
  check(response2, { "status is 404": (r) => r.status === 404 });
  sleep(.300);
};

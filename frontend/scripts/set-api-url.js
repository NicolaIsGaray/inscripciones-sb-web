/**
 * Genera src/environments/environment.prod.ts a partir de la variable
 * de entorno API_URL. Se ejecuta antes de `ng build` en el build de Render.
 *
 *   API_URL=https://mi-api.onrender.com/api npm run build
 *
 * Si API_URL no está definida, deja el archivo como está y avisa por consola.
 */
const fs = require('fs');
const path = require('path');

const target = path.join(__dirname, '..', 'src', 'environments', 'environment.prod.ts');
let apiUrl = (process.env.API_URL || '').trim().replace(/\/+$/, '');

if (!apiUrl) {
  console.warn(
    '[set-api-url] API_URL no definida: se mantiene el valor committed en environment.prod.ts. ' +
      'Definí API_URL en el panel de Render para apuntar al backend correcto.'
  );
  process.exit(0);
}

// Si trae esquema, que sea http/https
if (apiUrl.includes('://') && !/^https?:\/\//.test(apiUrl)) {
  console.error(`[set-api-url] API_URL inválida: "${apiUrl}". Debe empezar con http:// o https://`);
  process.exit(1);
}

// Render expone el hostname público sin esquema (ej. mi-api.onrender.com)
if (!/^https?:\/\//.test(apiUrl)) {
  apiUrl = `https://${apiUrl}`;
}

// El hostname de Render no trae path: la API vive bajo /api
if (!/^https?:\/\/[^/]+\/.+/.test(apiUrl)) {
  apiUrl = `${apiUrl}/api`;
}

// Validación final: esquema http(s), host no vacío, path opcional
if (!/^https?:\/\/[A-Za-z0-9](?:[A-Za-z0-9.\-]*[A-Za-z0-9])?(?::\d+)?(?:\/.*)?$/.test(apiUrl)) {
  console.error(`[set-api-url] API_URL inválida: "${apiUrl}". Debe ser algo como https://mi-api.onrender.com/api`);
  process.exit(1);
}

const contents = `/**
 * Configuración de PRODUCCIÓN.
 *
 * Generado por scripts/set-api-url.js — no editar a mano.
 * API_URL en tiempo de build: ${apiUrl}
 */
export const environment = {
  production: true,
  apiUrl: '${apiUrl}',
};
`;

fs.writeFileSync(target, contents, 'utf8');
console.log(`[set-api-url] environment.prod.ts apuntando a ${apiUrl}`);
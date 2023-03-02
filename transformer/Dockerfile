FROM container-runtime

ADD package.json package-lock.json ./
USER 0
RUN npm install
USER 1001

ADD transformer.js .


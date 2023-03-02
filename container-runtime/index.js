
const main = async () => {
    const buff = Buffer.from(process.env.DATA, 'base64');
    const data = JSON.parse(buff.toString())
    const transformer = await import('./app/transformer.js');
    console.log("About to execute the transformer")
    const result = await transformer.process(data);
    console.log(process.env.RESULT_NAME_ANCHOR +": "+JSON.stringify(result))
}

console.log("Executing runtime")
main()



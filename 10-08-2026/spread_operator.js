const customer = {
    name: "teja",
    age: 21,
    gender: "male",
    dateOfBirth: "2005-01-01",
    address: {
        street: "123 Main St",
        city: "Anytown",
        state: "CA",
        zip: "12345"
    }   
}
console.log(customer)
const updateCustomer = {
    ...customer,
    age: 22,
    address: {
        ...customer.address,
        city: "Newtown"
    }
}
console.log(updateCustomer)
object Solution {
    def defangIPaddr(address: String): String = {
        address.replace(".", "[.]")
    }
}
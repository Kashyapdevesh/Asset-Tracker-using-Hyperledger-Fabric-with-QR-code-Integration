var RefData = require("./RefData.json");
module.exports = {
    getOrgByUser(username) {
        var org = '';
        RefData.users.forEach(user => {
            if (user.username == username) {
                org = user.org;
            }
        });
        return org
    },
    getManufacturers() {
        return this.getOrgsByType('mfg')
    },
    getShippers() {
        return this.getOrgsByType('shp')
    },
    getRetailers() {
        return this.getOrgsByType('rtl')
    },
    getOrgsByType(typ) {
        var orgs = [];
        RefData.orgs.forEach(org => {
            if (org.orgType == typ) {
                orgs.push(org);
            }
        });
        return orgs
    }
    
}
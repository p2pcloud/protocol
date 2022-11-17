describe("BrokerV1_community", function () {
    describe("SetCommunityAddress", function () {
        it("should set community address");
        it("should revert if not owner");
    })
    describe("SetCommunityFee", function () {
        it("should set community fee");
        it("should revert if not owner");
        it("should revert if fee is not between 0 and 100");
    })
})
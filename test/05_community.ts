describe("BrokerV1_community", function () {
    describe("SetCommunityAddress", function () {
        it("should set community address");
        it("should revert if not owner");
    })
    describe("SetCommunityFee", function () {
        it("should set community fee");
        it("should change amount of fee paid in ClaimPament");
        it("should revert if not community contract");
        it("should revert if fee is not between 0 (0%) and 2500 (25%)");
    })
})
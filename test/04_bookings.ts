describe("BrokerV1_bookings", function () {
    describe("BookVM", function () {
        it("should create a new booking");
        it("should revert if not enough free stablecoin balance");
        it("should increase locked stablecoin balance");
        it("should decrease machines available");
        it("should revert if no machines available");
    })
    describe("TerminateVM", function () {
        it("should revert if booking does not exist");
        it("should revert if user is not the owner");
        it("should delete booking");
        it("should increase machines available");
        it("should execute payment");
    })
    describe("FindBookingsByUser", function () {
        it("should return array of bookings");
    })
    describe("FindBookingsByMiner", function () {
        it("should return array of bookings");
    })
    describe("GetBooking", function () {
        it("should return booking");
    })
    describe("ClaimPayment", function () {
        it("should revert if booking does not exist");
        it("should revert if user is not the miner");
        it("should transfer payment to miner and commision to community");
    })
})
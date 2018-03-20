pragma solidity 0.4.21;

contract EventTest {

	event TestEvent(string _a, string _b);

	function emitEvent() public {
		emit TestEvent("a", "b");
	}
}
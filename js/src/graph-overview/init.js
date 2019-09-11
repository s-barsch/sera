import yall from "./yall";
import observeMonths from "./months";
import initTopLink from "./top-link";
import initSubnavFix from "./subnav-fix";

export default function initOverview(pageType) {
    yall();
    if (pageType != "graph-main" || window.innerWidth > 1024) {
        initTopLink();
    }
    if (pageType == "graph-year") {
        initSubnavFix();
        observeMonths();
    };
}


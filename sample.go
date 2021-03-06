package main

const sampleData = `

<div class="outer-results-panel">
    <h2>Bin collection details for</h2>
    <h3></h3>
    <h3>Your next bin collections are:</h3>
    <div class="inner-results-contains-table">
        <div class="marginator">
            <table class="job-details">
                <tr class="spacer-row"><td class="graphic-cell"></td><td class="visible-cell"></td></tr>
                <tr class="data-row green-waste-bin-type" data-toggle="tooltip" title="Waste-SOUTH 4 REFUSE-070519">
                    <td class="graphic-cell"></td>
                    <td class="visible-cell">
                        <label for="Tuesday">Tuesday</label> &nbsp;
                        <label for="">28/05/2019</label> &nbsp;
                        <label for="Empty_Standard_Garden_Waste">Empty Standard Garden Waste</label>
                    </td>
                </tr>
                <tr class="spacer-row"><td></td></tr>
                <tr class="data-row recyclable-bin-type" data-toggle="tooltip" title="Waste-SOUTH 2 GARDEN-140519">
                    <td class="graphic-cell"></td>
                    <td class="visible-cell">
                        <label for="Tuesday">Tuesday</label> &nbsp;
                        <label for="">28/05/2019</label> &nbsp;
                        <label for="Empty_Standard_Mixed_Recycling">Empty Standard Mixed Recycling</label>
                    </td>
                </tr>
                <tr class="spacer-row"><td></td></tr>
                <tr class="data-row non-recyclable-bin-type" data-toggle="tooltip" title="Waste-SOUTH 3 SILVER-140519">
                    <td class="graphic-cell"></td>
                    <td class="visible-cell">
                        <label for="Tuesday">Tuesday</label> &nbsp;
                        <label for="">04/06/2019</label> &nbsp;
                        <label for="Empty_Standard_General_Waste">Empty Standard General Waste</label>
                    </td>
                </tr>
                <tr class="spacer-row"><td></td></tr>
                <tr class="data-row recyclable-bin-type" data-toggle="tooltip" title="Waste-SOUTH 4 REFUSE-210519">
                    <td class="graphic-cell"></td>
                    <td class="visible-cell">
                        <label for="Tuesday">Tuesday</label> &nbsp;
                        <label for="">11/06/2019</label> &nbsp;
                        <label for="Empty_Standard_Mixed_Recycling">Empty Standard Mixed Recycling</label>
                    </td>
                </tr>
                <tr class="spacer-row"><td></td></tr>
                <tr class="data-row green-waste-bin-type" data-toggle="tooltip" title="Waste-SOUTH 2 GARDEN-280519">
                    <td class="graphic-cell"></td>
                    <td class="visible-cell">
                        <label for="Tuesday">Tuesday</label> &nbsp;
                        <label for="">11/06/2019</label> &nbsp;
                        <label for="Empty_Standard_Garden_Waste">Empty Standard Garden Waste</label>
                    </td>
                </tr>
                <tr class="spacer-row"><td></td></tr>
                <tr class="data-row non-recyclable-bin-type" data-toggle="tooltip" title="Waste-SOUTH 3 SILVER-280519">
                    <td class="graphic-cell"></td>
                    <td class="visible-cell">
                        <label for="Tuesday">Tuesday</label> &nbsp;
                        <label for="">18/06/2019</label> &nbsp;
                        <label for="Empty_Standard_General_Waste">Empty Standard General Waste</label>
                    </td>
                </tr>
                <tr class="spacer-row"><td></td></tr>
                <tr class="data-row green-waste-bin-type" data-toggle="tooltip" title="Waste-SOUTH 4 REFUSE-040619">
                    <td class="graphic-cell"></td>
                    <td class="visible-cell">
                        <label for="Tuesday">Tuesday</label> &nbsp;
                        <label for="">25/06/2019</label> &nbsp;
                        <label for="Empty_Standard_Garden_Waste">Empty Standard Garden Waste</label>
                    </td>
                </tr>
                <tr class="spacer-row"><td></td></tr>
                <tr class="data-row recyclable-bin-type" data-toggle="tooltip" title="Waste-SOUTH 3 SILVER-110619">
                    <td class="graphic-cell"></td>
                    <td class="visible-cell">
                        <label for="Tuesday">Tuesday</label> &nbsp;
                        <label for="">25/06/2019</label> &nbsp;
                        <label for="Empty_Standard_Mixed_Recycling">Empty Standard Mixed Recycling</label>
                    </td>
                </tr>
                <tr class="spacer-row"><td></td></tr>
                <tr class="data-row non-recyclable-bin-type" data-toggle="tooltip" title="Waste-SOUTH 2 GARDEN-110619">
                    <td class="graphic-cell"></td>
                    <td class="visible-cell">
                        <label for="Tuesday">Tuesday</label> &nbsp;
                        <label for="">02/07/2019</label> &nbsp;
                        <label for="Empty_Standard_General_Waste">Empty Standard General Waste</label>
                    </td>
                </tr>
                <tr class="spacer-row"><td></td></tr>
                <tr class="data-row green-waste-bin-type" data-toggle="tooltip" title="Waste-SOUTH 4 REFUSE-180619">
                    <td class="graphic-cell"></td>
                    <td class="visible-cell">
                        <label for="Tuesday">Tuesday</label> &nbsp;
                        <label for="">09/07/2019</label> &nbsp;
                        <label for="Empty_Standard_Garden_Waste">Empty Standard Garden Waste</label>
                    </td>
                </tr>
                <tr class="spacer-row"><td></td></tr>

            </table>
        </div>
    </div>
    <br />
    <form action="/MyCollectionDay/SearchByAjax/PostCalendar" method="post">    <table class="full-job-list-for-posting">
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_0__Id" name="BartecSimplifiedJobList[0].Id" type="hidden" value="107053925" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_0__Name" name="BartecSimplifiedJobList[0].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_0__ScheduledStart" name="BartecSimplifiedJobList[0].ScheduledStart" type="hidden" value="07/05/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_0__JobDescription" name="BartecSimplifiedJobList[0].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_0__IsRescheduled" name="BartecSimplifiedJobList[0].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_1__Id" name="BartecSimplifiedJobList[1].Id" type="hidden" value="98129042" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_1__Name" name="BartecSimplifiedJobList[1].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_1__ScheduledStart" name="BartecSimplifiedJobList[1].ScheduledStart" type="hidden" value="14/05/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_1__JobDescription" name="BartecSimplifiedJobList[1].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_1__IsRescheduled" name="BartecSimplifiedJobList[1].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_2__Id" name="BartecSimplifiedJobList[2].Id" type="hidden" value="107264278" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_2__Name" name="BartecSimplifiedJobList[2].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_2__ScheduledStart" name="BartecSimplifiedJobList[2].ScheduledStart" type="hidden" value="14/05/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_2__JobDescription" name="BartecSimplifiedJobList[2].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_2__IsRescheduled" name="BartecSimplifiedJobList[2].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_3__Id" name="BartecSimplifiedJobList[3].Id" type="hidden" value="107425213" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_3__Name" name="BartecSimplifiedJobList[3].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_3__ScheduledStart" name="BartecSimplifiedJobList[3].ScheduledStart" type="hidden" value="21/05/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_3__JobDescription" name="BartecSimplifiedJobList[3].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_3__IsRescheduled" name="BartecSimplifiedJobList[3].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_4__Id" name="BartecSimplifiedJobList[4].Id" type="hidden" value="98262423" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_4__Name" name="BartecSimplifiedJobList[4].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_4__ScheduledStart" name="BartecSimplifiedJobList[4].ScheduledStart" type="hidden" value="28/05/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_4__JobDescription" name="BartecSimplifiedJobList[4].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_4__IsRescheduled" name="BartecSimplifiedJobList[4].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_5__Id" name="BartecSimplifiedJobList[5].Id" type="hidden" value="107641181" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_5__Name" name="BartecSimplifiedJobList[5].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_5__ScheduledStart" name="BartecSimplifiedJobList[5].ScheduledStart" type="hidden" value="28/05/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_5__JobDescription" name="BartecSimplifiedJobList[5].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_5__IsRescheduled" name="BartecSimplifiedJobList[5].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_6__Id" name="BartecSimplifiedJobList[6].Id" type="hidden" value="107798939" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_6__Name" name="BartecSimplifiedJobList[6].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_6__ScheduledStart" name="BartecSimplifiedJobList[6].ScheduledStart" type="hidden" value="04/06/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_6__JobDescription" name="BartecSimplifiedJobList[6].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_6__IsRescheduled" name="BartecSimplifiedJobList[6].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_7__Id" name="BartecSimplifiedJobList[7].Id" type="hidden" value="108004208" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_7__Name" name="BartecSimplifiedJobList[7].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_7__ScheduledStart" name="BartecSimplifiedJobList[7].ScheduledStart" type="hidden" value="11/06/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_7__JobDescription" name="BartecSimplifiedJobList[7].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_7__IsRescheduled" name="BartecSimplifiedJobList[7].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_8__Id" name="BartecSimplifiedJobList[8].Id" type="hidden" value="98694252" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_8__Name" name="BartecSimplifiedJobList[8].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_8__ScheduledStart" name="BartecSimplifiedJobList[8].ScheduledStart" type="hidden" value="11/06/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_8__JobDescription" name="BartecSimplifiedJobList[8].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_8__IsRescheduled" name="BartecSimplifiedJobList[8].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_9__Id" name="BartecSimplifiedJobList[9].Id" type="hidden" value="108158682" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_9__Name" name="BartecSimplifiedJobList[9].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_9__ScheduledStart" name="BartecSimplifiedJobList[9].ScheduledStart" type="hidden" value="18/06/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_9__JobDescription" name="BartecSimplifiedJobList[9].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_9__IsRescheduled" name="BartecSimplifiedJobList[9].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_10__Id" name="BartecSimplifiedJobList[10].Id" type="hidden" value="98864604" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_10__Name" name="BartecSimplifiedJobList[10].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_10__ScheduledStart" name="BartecSimplifiedJobList[10].ScheduledStart" type="hidden" value="25/06/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_10__JobDescription" name="BartecSimplifiedJobList[10].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_10__IsRescheduled" name="BartecSimplifiedJobList[10].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_11__Id" name="BartecSimplifiedJobList[11].Id" type="hidden" value="108379590" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_11__Name" name="BartecSimplifiedJobList[11].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_11__ScheduledStart" name="BartecSimplifiedJobList[11].ScheduledStart" type="hidden" value="25/06/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_11__JobDescription" name="BartecSimplifiedJobList[11].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_11__IsRescheduled" name="BartecSimplifiedJobList[11].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_12__Id" name="BartecSimplifiedJobList[12].Id" type="hidden" value="108549341" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_12__Name" name="BartecSimplifiedJobList[12].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_12__ScheduledStart" name="BartecSimplifiedJobList[12].ScheduledStart" type="hidden" value="02/07/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_12__JobDescription" name="BartecSimplifiedJobList[12].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_12__IsRescheduled" name="BartecSimplifiedJobList[12].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_13__Id" name="BartecSimplifiedJobList[13].Id" type="hidden" value="99039376" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_13__Name" name="BartecSimplifiedJobList[13].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_13__ScheduledStart" name="BartecSimplifiedJobList[13].ScheduledStart" type="hidden" value="09/07/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_13__JobDescription" name="BartecSimplifiedJobList[13].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_13__IsRescheduled" name="BartecSimplifiedJobList[13].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_14__Id" name="BartecSimplifiedJobList[14].Id" type="hidden" value="108767321" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_14__Name" name="BartecSimplifiedJobList[14].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_14__ScheduledStart" name="BartecSimplifiedJobList[14].ScheduledStart" type="hidden" value="09/07/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_14__JobDescription" name="BartecSimplifiedJobList[14].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_14__IsRescheduled" name="BartecSimplifiedJobList[14].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_15__Id" name="BartecSimplifiedJobList[15].Id" type="hidden" value="108913961" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_15__Name" name="BartecSimplifiedJobList[15].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_15__ScheduledStart" name="BartecSimplifiedJobList[15].ScheduledStart" type="hidden" value="16/07/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_15__JobDescription" name="BartecSimplifiedJobList[15].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_15__IsRescheduled" name="BartecSimplifiedJobList[15].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_16__Id" name="BartecSimplifiedJobList[16].Id" type="hidden" value="98976984" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_16__Name" name="BartecSimplifiedJobList[16].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_16__ScheduledStart" name="BartecSimplifiedJobList[16].ScheduledStart" type="hidden" value="23/07/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_16__JobDescription" name="BartecSimplifiedJobList[16].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_16__IsRescheduled" name="BartecSimplifiedJobList[16].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_17__Id" name="BartecSimplifiedJobList[17].Id" type="hidden" value="109126067" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_17__Name" name="BartecSimplifiedJobList[17].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_17__ScheduledStart" name="BartecSimplifiedJobList[17].ScheduledStart" type="hidden" value="23/07/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_17__JobDescription" name="BartecSimplifiedJobList[17].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_17__IsRescheduled" name="BartecSimplifiedJobList[17].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_18__Id" name="BartecSimplifiedJobList[18].Id" type="hidden" value="109272207" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_18__Name" name="BartecSimplifiedJobList[18].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_18__ScheduledStart" name="BartecSimplifiedJobList[18].ScheduledStart" type="hidden" value="30/07/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_18__JobDescription" name="BartecSimplifiedJobList[18].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_18__IsRescheduled" name="BartecSimplifiedJobList[18].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_19__Id" name="BartecSimplifiedJobList[19].Id" type="hidden" value="109489688" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_19__Name" name="BartecSimplifiedJobList[19].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_19__ScheduledStart" name="BartecSimplifiedJobList[19].ScheduledStart" type="hidden" value="06/08/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_19__JobDescription" name="BartecSimplifiedJobList[19].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_19__IsRescheduled" name="BartecSimplifiedJobList[19].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_20__Id" name="BartecSimplifiedJobList[20].Id" type="hidden" value="99207362" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_20__Name" name="BartecSimplifiedJobList[20].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_20__ScheduledStart" name="BartecSimplifiedJobList[20].ScheduledStart" type="hidden" value="06/08/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_20__JobDescription" name="BartecSimplifiedJobList[20].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_20__IsRescheduled" name="BartecSimplifiedJobList[20].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_21__Id" name="BartecSimplifiedJobList[21].Id" type="hidden" value="109660703" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_21__Name" name="BartecSimplifiedJobList[21].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_21__ScheduledStart" name="BartecSimplifiedJobList[21].ScheduledStart" type="hidden" value="13/08/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_21__JobDescription" name="BartecSimplifiedJobList[21].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_21__IsRescheduled" name="BartecSimplifiedJobList[21].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_22__Id" name="BartecSimplifiedJobList[22].Id" type="hidden" value="99351278" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_22__Name" name="BartecSimplifiedJobList[22].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_22__ScheduledStart" name="BartecSimplifiedJobList[22].ScheduledStart" type="hidden" value="20/08/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_22__JobDescription" name="BartecSimplifiedJobList[22].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_22__IsRescheduled" name="BartecSimplifiedJobList[22].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_23__Id" name="BartecSimplifiedJobList[23].Id" type="hidden" value="109876953" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_23__Name" name="BartecSimplifiedJobList[23].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_23__ScheduledStart" name="BartecSimplifiedJobList[23].ScheduledStart" type="hidden" value="20/08/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_23__JobDescription" name="BartecSimplifiedJobList[23].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_23__IsRescheduled" name="BartecSimplifiedJobList[23].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_24__Id" name="BartecSimplifiedJobList[24].Id" type="hidden" value="110021749" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_24__Name" name="BartecSimplifiedJobList[24].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_24__ScheduledStart" name="BartecSimplifiedJobList[24].ScheduledStart" type="hidden" value="27/08/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_24__JobDescription" name="BartecSimplifiedJobList[24].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_24__IsRescheduled" name="BartecSimplifiedJobList[24].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_25__Id" name="BartecSimplifiedJobList[25].Id" type="hidden" value="110231559" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_25__Name" name="BartecSimplifiedJobList[25].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_25__ScheduledStart" name="BartecSimplifiedJobList[25].ScheduledStart" type="hidden" value="03/09/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_25__JobDescription" name="BartecSimplifiedJobList[25].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_25__IsRescheduled" name="BartecSimplifiedJobList[25].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_26__Id" name="BartecSimplifiedJobList[26].Id" type="hidden" value="99532313" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_26__Name" name="BartecSimplifiedJobList[26].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_26__ScheduledStart" name="BartecSimplifiedJobList[26].ScheduledStart" type="hidden" value="03/09/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_26__JobDescription" name="BartecSimplifiedJobList[26].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_26__IsRescheduled" name="BartecSimplifiedJobList[26].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_27__Id" name="BartecSimplifiedJobList[27].Id" type="hidden" value="110407036" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_27__Name" name="BartecSimplifiedJobList[27].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_27__ScheduledStart" name="BartecSimplifiedJobList[27].ScheduledStart" type="hidden" value="10/09/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_27__JobDescription" name="BartecSimplifiedJobList[27].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_27__IsRescheduled" name="BartecSimplifiedJobList[27].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_28__Id" name="BartecSimplifiedJobList[28].Id" type="hidden" value="110602998" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_28__Name" name="BartecSimplifiedJobList[28].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_28__ScheduledStart" name="BartecSimplifiedJobList[28].ScheduledStart" type="hidden" value="17/09/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_28__JobDescription" name="BartecSimplifiedJobList[28].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_28__IsRescheduled" name="BartecSimplifiedJobList[28].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_29__Id" name="BartecSimplifiedJobList[29].Id" type="hidden" value="99719769" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_29__Name" name="BartecSimplifiedJobList[29].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_29__ScheduledStart" name="BartecSimplifiedJobList[29].ScheduledStart" type="hidden" value="17/09/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_29__JobDescription" name="BartecSimplifiedJobList[29].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_29__IsRescheduled" name="BartecSimplifiedJobList[29].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_30__Id" name="BartecSimplifiedJobList[30].Id" type="hidden" value="110777967" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_30__Name" name="BartecSimplifiedJobList[30].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_30__ScheduledStart" name="BartecSimplifiedJobList[30].ScheduledStart" type="hidden" value="24/09/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_30__JobDescription" name="BartecSimplifiedJobList[30].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_30__IsRescheduled" name="BartecSimplifiedJobList[30].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_31__Id" name="BartecSimplifiedJobList[31].Id" type="hidden" value="110982026" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_31__Name" name="BartecSimplifiedJobList[31].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_31__ScheduledStart" name="BartecSimplifiedJobList[31].ScheduledStart" type="hidden" value="01/10/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_31__JobDescription" name="BartecSimplifiedJobList[31].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_31__IsRescheduled" name="BartecSimplifiedJobList[31].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_32__Id" name="BartecSimplifiedJobList[32].Id" type="hidden" value="99911647" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_32__Name" name="BartecSimplifiedJobList[32].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_32__ScheduledStart" name="BartecSimplifiedJobList[32].ScheduledStart" type="hidden" value="01/10/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_32__JobDescription" name="BartecSimplifiedJobList[32].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_32__IsRescheduled" name="BartecSimplifiedJobList[32].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_33__Id" name="BartecSimplifiedJobList[33].Id" type="hidden" value="111157559" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_33__Name" name="BartecSimplifiedJobList[33].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_33__ScheduledStart" name="BartecSimplifiedJobList[33].ScheduledStart" type="hidden" value="08/10/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_33__JobDescription" name="BartecSimplifiedJobList[33].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_33__IsRescheduled" name="BartecSimplifiedJobList[33].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_34__Id" name="BartecSimplifiedJobList[34].Id" type="hidden" value="100112835" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_34__Name" name="BartecSimplifiedJobList[34].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_34__ScheduledStart" name="BartecSimplifiedJobList[34].ScheduledStart" type="hidden" value="15/10/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_34__JobDescription" name="BartecSimplifiedJobList[34].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_34__IsRescheduled" name="BartecSimplifiedJobList[34].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_35__Id" name="BartecSimplifiedJobList[35].Id" type="hidden" value="111336874" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_35__Name" name="BartecSimplifiedJobList[35].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_35__ScheduledStart" name="BartecSimplifiedJobList[35].ScheduledStart" type="hidden" value="15/10/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_35__JobDescription" name="BartecSimplifiedJobList[35].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_35__IsRescheduled" name="BartecSimplifiedJobList[35].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_36__Id" name="BartecSimplifiedJobList[36].Id" type="hidden" value="111514469" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_36__Name" name="BartecSimplifiedJobList[36].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_36__ScheduledStart" name="BartecSimplifiedJobList[36].ScheduledStart" type="hidden" value="22/10/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_36__JobDescription" name="BartecSimplifiedJobList[36].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_36__IsRescheduled" name="BartecSimplifiedJobList[36].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_37__Id" name="BartecSimplifiedJobList[37].Id" type="hidden" value="111729263" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_37__Name" name="BartecSimplifiedJobList[37].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_37__ScheduledStart" name="BartecSimplifiedJobList[37].ScheduledStart" type="hidden" value="29/10/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_37__JobDescription" name="BartecSimplifiedJobList[37].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_37__IsRescheduled" name="BartecSimplifiedJobList[37].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_38__Id" name="BartecSimplifiedJobList[38].Id" type="hidden" value="100438426" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_38__Name" name="BartecSimplifiedJobList[38].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_38__ScheduledStart" name="BartecSimplifiedJobList[38].ScheduledStart" type="hidden" value="29/10/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_38__JobDescription" name="BartecSimplifiedJobList[38].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_38__IsRescheduled" name="BartecSimplifiedJobList[38].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_39__Id" name="BartecSimplifiedJobList[39].Id" type="hidden" value="111902159" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_39__Name" name="BartecSimplifiedJobList[39].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_39__ScheduledStart" name="BartecSimplifiedJobList[39].ScheduledStart" type="hidden" value="05/11/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_39__JobDescription" name="BartecSimplifiedJobList[39].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_39__IsRescheduled" name="BartecSimplifiedJobList[39].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_40__Id" name="BartecSimplifiedJobList[40].Id" type="hidden" value="100998561" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_40__Name" name="BartecSimplifiedJobList[40].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_40__ScheduledStart" name="BartecSimplifiedJobList[40].ScheduledStart" type="hidden" value="12/11/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_40__JobDescription" name="BartecSimplifiedJobList[40].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_40__IsRescheduled" name="BartecSimplifiedJobList[40].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_41__Id" name="BartecSimplifiedJobList[41].Id" type="hidden" value="112122086" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_41__Name" name="BartecSimplifiedJobList[41].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_41__ScheduledStart" name="BartecSimplifiedJobList[41].ScheduledStart" type="hidden" value="12/11/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_41__JobDescription" name="BartecSimplifiedJobList[41].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_41__IsRescheduled" name="BartecSimplifiedJobList[41].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_42__Id" name="BartecSimplifiedJobList[42].Id" type="hidden" value="112274084" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_42__Name" name="BartecSimplifiedJobList[42].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_42__ScheduledStart" name="BartecSimplifiedJobList[42].ScheduledStart" type="hidden" value="19/11/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_42__JobDescription" name="BartecSimplifiedJobList[42].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_42__IsRescheduled" name="BartecSimplifiedJobList[42].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_43__Id" name="BartecSimplifiedJobList[43].Id" type="hidden" value="101560452" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_43__Name" name="BartecSimplifiedJobList[43].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_43__ScheduledStart" name="BartecSimplifiedJobList[43].ScheduledStart" type="hidden" value="26/11/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_43__JobDescription" name="BartecSimplifiedJobList[43].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_43__IsRescheduled" name="BartecSimplifiedJobList[43].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_44__Id" name="BartecSimplifiedJobList[44].Id" type="hidden" value="112476291" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_44__Name" name="BartecSimplifiedJobList[44].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_44__ScheduledStart" name="BartecSimplifiedJobList[44].ScheduledStart" type="hidden" value="26/11/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_44__JobDescription" name="BartecSimplifiedJobList[44].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_44__IsRescheduled" name="BartecSimplifiedJobList[44].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_45__Id" name="BartecSimplifiedJobList[45].Id" type="hidden" value="112629800" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_45__Name" name="BartecSimplifiedJobList[45].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_45__ScheduledStart" name="BartecSimplifiedJobList[45].ScheduledStart" type="hidden" value="03/12/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_45__JobDescription" name="BartecSimplifiedJobList[45].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_45__IsRescheduled" name="BartecSimplifiedJobList[45].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_46__Id" name="BartecSimplifiedJobList[46].Id" type="hidden" value="112838966" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_46__Name" name="BartecSimplifiedJobList[46].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_46__ScheduledStart" name="BartecSimplifiedJobList[46].ScheduledStart" type="hidden" value="10/12/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_46__JobDescription" name="BartecSimplifiedJobList[46].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_46__IsRescheduled" name="BartecSimplifiedJobList[46].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_47__Id" name="BartecSimplifiedJobList[47].Id" type="hidden" value="102127308" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_47__Name" name="BartecSimplifiedJobList[47].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_47__ScheduledStart" name="BartecSimplifiedJobList[47].ScheduledStart" type="hidden" value="10/12/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_47__JobDescription" name="BartecSimplifiedJobList[47].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_47__IsRescheduled" name="BartecSimplifiedJobList[47].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_48__Id" name="BartecSimplifiedJobList[48].Id" type="hidden" value="113011619" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_48__Name" name="BartecSimplifiedJobList[48].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_48__ScheduledStart" name="BartecSimplifiedJobList[48].ScheduledStart" type="hidden" value="17/12/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_48__JobDescription" name="BartecSimplifiedJobList[48].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_48__IsRescheduled" name="BartecSimplifiedJobList[48].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_49__Id" name="BartecSimplifiedJobList[49].Id" type="hidden" value="102721892" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_49__Name" name="BartecSimplifiedJobList[49].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_49__ScheduledStart" name="BartecSimplifiedJobList[49].ScheduledStart" type="hidden" value="24/12/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_49__JobDescription" name="BartecSimplifiedJobList[49].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_49__IsRescheduled" name="BartecSimplifiedJobList[49].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_50__Id" name="BartecSimplifiedJobList[50].Id" type="hidden" value="113217362" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_50__Name" name="BartecSimplifiedJobList[50].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_50__ScheduledStart" name="BartecSimplifiedJobList[50].ScheduledStart" type="hidden" value="24/12/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_50__JobDescription" name="BartecSimplifiedJobList[50].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_50__IsRescheduled" name="BartecSimplifiedJobList[50].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_51__Id" name="BartecSimplifiedJobList[51].Id" type="hidden" value="113436471" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_51__Name" name="BartecSimplifiedJobList[51].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_51__ScheduledStart" name="BartecSimplifiedJobList[51].ScheduledStart" type="hidden" value="31/12/2019 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_51__JobDescription" name="BartecSimplifiedJobList[51].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_51__IsRescheduled" name="BartecSimplifiedJobList[51].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_52__Id" name="BartecSimplifiedJobList[52].Id" type="hidden" value="113773192" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_52__Name" name="BartecSimplifiedJobList[52].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_52__ScheduledStart" name="BartecSimplifiedJobList[52].ScheduledStart" type="hidden" value="07/01/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_52__JobDescription" name="BartecSimplifiedJobList[52].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_52__IsRescheduled" name="BartecSimplifiedJobList[52].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_53__Id" name="BartecSimplifiedJobList[53].Id" type="hidden" value="113812292" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_53__Name" name="BartecSimplifiedJobList[53].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_53__ScheduledStart" name="BartecSimplifiedJobList[53].ScheduledStart" type="hidden" value="07/01/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_53__JobDescription" name="BartecSimplifiedJobList[53].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_53__IsRescheduled" name="BartecSimplifiedJobList[53].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_54__Id" name="BartecSimplifiedJobList[54].Id" type="hidden" value="114125082" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_54__Name" name="BartecSimplifiedJobList[54].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_54__ScheduledStart" name="BartecSimplifiedJobList[54].ScheduledStart" type="hidden" value="14/01/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_54__JobDescription" name="BartecSimplifiedJobList[54].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_54__IsRescheduled" name="BartecSimplifiedJobList[54].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_55__Id" name="BartecSimplifiedJobList[55].Id" type="hidden" value="114201373" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_55__Name" name="BartecSimplifiedJobList[55].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_55__ScheduledStart" name="BartecSimplifiedJobList[55].ScheduledStart" type="hidden" value="21/01/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_55__JobDescription" name="BartecSimplifiedJobList[55].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_55__IsRescheduled" name="BartecSimplifiedJobList[55].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_56__Id" name="BartecSimplifiedJobList[56].Id" type="hidden" value="114364954" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_56__Name" name="BartecSimplifiedJobList[56].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_56__ScheduledStart" name="BartecSimplifiedJobList[56].ScheduledStart" type="hidden" value="21/01/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_56__JobDescription" name="BartecSimplifiedJobList[56].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_56__IsRescheduled" name="BartecSimplifiedJobList[56].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_57__Id" name="BartecSimplifiedJobList[57].Id" type="hidden" value="114523507" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_57__Name" name="BartecSimplifiedJobList[57].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_57__ScheduledStart" name="BartecSimplifiedJobList[57].ScheduledStart" type="hidden" value="28/01/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_57__JobDescription" name="BartecSimplifiedJobList[57].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_57__IsRescheduled" name="BartecSimplifiedJobList[57].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_58__Id" name="BartecSimplifiedJobList[58].Id" type="hidden" value="114953698" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_58__Name" name="BartecSimplifiedJobList[58].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_58__ScheduledStart" name="BartecSimplifiedJobList[58].ScheduledStart" type="hidden" value="04/02/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_58__JobDescription" name="BartecSimplifiedJobList[58].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_58__IsRescheduled" name="BartecSimplifiedJobList[58].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_59__Id" name="BartecSimplifiedJobList[59].Id" type="hidden" value="114972390" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_59__Name" name="BartecSimplifiedJobList[59].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_59__ScheduledStart" name="BartecSimplifiedJobList[59].ScheduledStart" type="hidden" value="04/02/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_59__JobDescription" name="BartecSimplifiedJobList[59].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_59__IsRescheduled" name="BartecSimplifiedJobList[59].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_60__Id" name="BartecSimplifiedJobList[60].Id" type="hidden" value="115280589" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_60__Name" name="BartecSimplifiedJobList[60].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_60__ScheduledStart" name="BartecSimplifiedJobList[60].ScheduledStart" type="hidden" value="11/02/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_60__JobDescription" name="BartecSimplifiedJobList[60].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_60__IsRescheduled" name="BartecSimplifiedJobList[60].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_61__Id" name="BartecSimplifiedJobList[61].Id" type="hidden" value="115530221" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_61__Name" name="BartecSimplifiedJobList[61].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_61__ScheduledStart" name="BartecSimplifiedJobList[61].ScheduledStart" type="hidden" value="18/02/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_61__JobDescription" name="BartecSimplifiedJobList[61].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_61__IsRescheduled" name="BartecSimplifiedJobList[61].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_62__Id" name="BartecSimplifiedJobList[62].Id" type="hidden" value="115557572" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_62__Name" name="BartecSimplifiedJobList[62].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_62__ScheduledStart" name="BartecSimplifiedJobList[62].ScheduledStart" type="hidden" value="18/02/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_62__JobDescription" name="BartecSimplifiedJobList[62].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_62__IsRescheduled" name="BartecSimplifiedJobList[62].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_63__Id" name="BartecSimplifiedJobList[63].Id" type="hidden" value="115885782" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_63__Name" name="BartecSimplifiedJobList[63].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_63__ScheduledStart" name="BartecSimplifiedJobList[63].ScheduledStart" type="hidden" value="25/02/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_63__JobDescription" name="BartecSimplifiedJobList[63].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_63__IsRescheduled" name="BartecSimplifiedJobList[63].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_64__Id" name="BartecSimplifiedJobList[64].Id" type="hidden" value="116122301" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_64__Name" name="BartecSimplifiedJobList[64].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_64__ScheduledStart" name="BartecSimplifiedJobList[64].ScheduledStart" type="hidden" value="03/03/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_64__JobDescription" name="BartecSimplifiedJobList[64].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_64__IsRescheduled" name="BartecSimplifiedJobList[64].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_65__Id" name="BartecSimplifiedJobList[65].Id" type="hidden" value="116283473" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_65__Name" name="BartecSimplifiedJobList[65].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_65__ScheduledStart" name="BartecSimplifiedJobList[65].ScheduledStart" type="hidden" value="03/03/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_65__JobDescription" name="BartecSimplifiedJobList[65].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_65__IsRescheduled" name="BartecSimplifiedJobList[65].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_66__Id" name="BartecSimplifiedJobList[66].Id" type="hidden" value="116631944" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_66__Name" name="BartecSimplifiedJobList[66].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_66__ScheduledStart" name="BartecSimplifiedJobList[66].ScheduledStart" type="hidden" value="10/03/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_66__JobDescription" name="BartecSimplifiedJobList[66].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_66__IsRescheduled" name="BartecSimplifiedJobList[66].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_67__Id" name="BartecSimplifiedJobList[67].Id" type="hidden" value="116788628" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_67__Name" name="BartecSimplifiedJobList[67].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_67__ScheduledStart" name="BartecSimplifiedJobList[67].ScheduledStart" type="hidden" value="17/03/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_67__JobDescription" name="BartecSimplifiedJobList[67].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_67__IsRescheduled" name="BartecSimplifiedJobList[67].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_68__Id" name="BartecSimplifiedJobList[68].Id" type="hidden" value="116949137" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_68__Name" name="BartecSimplifiedJobList[68].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_68__ScheduledStart" name="BartecSimplifiedJobList[68].ScheduledStart" type="hidden" value="17/03/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_68__JobDescription" name="BartecSimplifiedJobList[68].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_68__IsRescheduled" name="BartecSimplifiedJobList[68].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_69__Id" name="BartecSimplifiedJobList[69].Id" type="hidden" value="117190681" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_69__Name" name="BartecSimplifiedJobList[69].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_69__ScheduledStart" name="BartecSimplifiedJobList[69].ScheduledStart" type="hidden" value="24/03/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_69__JobDescription" name="BartecSimplifiedJobList[69].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_69__IsRescheduled" name="BartecSimplifiedJobList[69].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_70__Id" name="BartecSimplifiedJobList[70].Id" type="hidden" value="117006631" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_70__Name" name="BartecSimplifiedJobList[70].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_70__ScheduledStart" name="BartecSimplifiedJobList[70].ScheduledStart" type="hidden" value="31/03/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_70__JobDescription" name="BartecSimplifiedJobList[70].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_70__IsRescheduled" name="BartecSimplifiedJobList[70].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_71__Id" name="BartecSimplifiedJobList[71].Id" type="hidden" value="117261795" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_71__Name" name="BartecSimplifiedJobList[71].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_71__ScheduledStart" name="BartecSimplifiedJobList[71].ScheduledStart" type="hidden" value="31/03/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_71__JobDescription" name="BartecSimplifiedJobList[71].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_71__IsRescheduled" name="BartecSimplifiedJobList[71].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_72__Id" name="BartecSimplifiedJobList[72].Id" type="hidden" value="117608921" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_72__Name" name="BartecSimplifiedJobList[72].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_72__ScheduledStart" name="BartecSimplifiedJobList[72].ScheduledStart" type="hidden" value="07/04/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_72__JobDescription" name="BartecSimplifiedJobList[72].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_72__IsRescheduled" name="BartecSimplifiedJobList[72].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_73__Id" name="BartecSimplifiedJobList[73].Id" type="hidden" value="117781031" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_73__Name" name="BartecSimplifiedJobList[73].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_73__ScheduledStart" name="BartecSimplifiedJobList[73].ScheduledStart" type="hidden" value="14/04/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_73__JobDescription" name="BartecSimplifiedJobList[73].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_73__IsRescheduled" name="BartecSimplifiedJobList[73].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_74__Id" name="BartecSimplifiedJobList[74].Id" type="hidden" value="117909351" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_74__Name" name="BartecSimplifiedJobList[74].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_74__ScheduledStart" name="BartecSimplifiedJobList[74].ScheduledStart" type="hidden" value="14/04/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_74__JobDescription" name="BartecSimplifiedJobList[74].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_74__IsRescheduled" name="BartecSimplifiedJobList[74].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row non-recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_75__Id" name="BartecSimplifiedJobList[75].Id" type="hidden" value="118107755" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_75__Name" name="BartecSimplifiedJobList[75].Name" type="hidden" value="Empty Standard General Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_75__ScheduledStart" name="BartecSimplifiedJobList[75].ScheduledStart" type="hidden" value="21/04/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_75__JobDescription" name="BartecSimplifiedJobList[75].JobDescription" type="hidden" value="Empty 240L BLACK" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_75__IsRescheduled" name="BartecSimplifiedJobList[75].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row recyclable-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_76__Id" name="BartecSimplifiedJobList[76].Id" type="hidden" value="118320900" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_76__Name" name="BartecSimplifiedJobList[76].Name" type="hidden" value="Empty Standard Mixed Recycling" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_76__ScheduledStart" name="BartecSimplifiedJobList[76].ScheduledStart" type="hidden" value="28/04/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_76__JobDescription" name="BartecSimplifiedJobList[76].JobDescription" type="hidden" value="Empty 240L SILVER" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_76__IsRescheduled" name="BartecSimplifiedJobList[76].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
            <tr class="data-row green-waste-bin-type">
                <td class="hidden-cell">
                    <input data-val="true" data-val-number="The field Id must be a number." data-val-required="The Id field is required." id="BartecSimplifiedJobList_77__Id" name="BartecSimplifiedJobList[77].Id" type="hidden" value="118367194" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_77__Name" name="BartecSimplifiedJobList[77].Name" type="hidden" value="Empty Standard Garden Waste" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-date="The field ScheduledStart must be a date." data-val-required="The ScheduledStart field is required." id="BartecSimplifiedJobList_77__ScheduledStart" name="BartecSimplifiedJobList[77].ScheduledStart" type="hidden" value="28/04/2020 07:00:00" />
                </td>
                <td class="hidden-cell">
                    <input id="BartecSimplifiedJobList_77__JobDescription" name="BartecSimplifiedJobList[77].JobDescription" type="hidden" value="Empty 240L GREEN" />
                </td>
                <td class="hidden-cell">
                    <input data-val="true" data-val-required="The IsRescheduled field is required." id="BartecSimplifiedJobList_77__IsRescheduled" name="BartecSimplifiedJobList[77].IsRescheduled" type="hidden" value="False" />
                </td>
            </tr>
        </table>
        <input style="display:none" type="text" name="OnelineAddress" />
        <div class="marginator">
            <button type="submit" class="btn job-calendar" tabindex="11">Download your bin collection calendar (PDF)</button>
        </div>
    </form>    <div class="centralising-panel">
        <form action="/MyCollectionDay/" method="post">            <button type="submit" class="btn btn-primary" tabindex="12">Cancel</button>
        </form>    </div>
    <br />
    <div class="container-fluid">
        <table class="table-responsive">
            <tr>
                <td>Non-recyclable</td>
                <td>Recycling</td>
                <td>Garden</td>
            </tr>
            <tr>
                <td><label for="SOUTH_4_REFUSE_Tue_Wk_B">SOUTH 4 REFUSE Tue Wk B</label></td>
                <td><label for="SOUTH_3_SILVER_Tue_Wk_A">SOUTH 3 SILVER Tue Wk A</label></td>
                <td><label for="SOUTH_2_GARDEN_Tue_Wk_A">SOUTH 2 GARDEN Tue Wk A</label></td>
            </tr>
        </table>
    </div>
</div>
<script type="text/javascript">
    jQuery(function () {

        // set focus on first input control
        $(".job-calendar").focus();
    });
</script>

`